package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
	"strconv"
)

type ReleaseMessageService interface {
	//正常发布和灰度发布
	Create(release *models2.ReleaseRequest) error
	ReleaseGrayTotal(namespaceId, name, appId, cluster, operator string, isDeleted bool) error
	Creates(releaseRequests []*models2.ReleaseRequest) error //批量发布
}

type releaseMessageService struct {
	repository               repositories.ReleaseMessageRepository
	releaseRepository        repositories.Release
	appNamespaceRepository   repositories.AppNamespaceRepository
	itemRepository           repositories.ItemRepisitory
	releaseHistoryRepository repositories.ReleaseHistoryRepository
	db                       *gorm.DB
}

func NewReleaseMessageService(
	repository repositories.ReleaseMessageRepository,
	releaseRepository repositories.Release,
	appNamespaceRepository repositories.AppNamespaceRepository,
	releaseHistoryRepository repositories.ReleaseHistoryRepository,
	itemRepository repositories.ItemRepisitory,
	db *gorm.DB,
) ReleaseMessageService {
	return &releaseMessageService{
		repository:               repository,
		releaseRepository:        releaseRepository,
		appNamespaceRepository:   appNamespaceRepository,
		itemRepository:           itemRepository,
		releaseHistoryRepository: releaseHistoryRepository,
		db:                       db,
	}
}

func (s releaseMessageService) Creates(releaseRequests []*models2.ReleaseRequest) error {
	for _, r := range releaseRequests {
		if err := s.Create(r); err != nil {
			return errors.Wrap(err, "release stop:"+r.AppId)
		}
	}
	return nil
}

//发布流程，查询需要发布的key，修改对应状态，删除以前发布配置，发布此次配置，发布最新版本号通知客户端，删除以往版本号，记录发布历史
func (s releaseMessageService) Create(releaseRequest *models2.ReleaseRequest) error {
	release := new(models.Release)
	if releaseRequest.AppId == "" {
		app, err := s.appNamespaceRepository.FindAppNamespaceById(strconv.Itoa(int(releaseRequest.NamespaceId)))
		if err != nil {
			return errors.Wrap(err, "call appNamespaceRepository.FindAppNamespaceById() error")
		}
		releaseRequest.AppId = app.AppId
		releaseRequest.ClusterName = app.ClusterName
		releaseRequest.LaneName = app.LaneName
		releaseRequest.Name = app.Name
	}
	release.DataChange_LastModifiedBy = releaseRequest.Operator
	release.DataChange_CreatedBy = releaseRequest.Operator
	release.NamespaceName = releaseRequest.Name
	release.AppId = releaseRequest.AppId
	release.Comment = releaseRequest.Comment
	release.LaneName = releaseRequest.LaneName
	release.ClusterName = releaseRequest.ClusterName
	release.ReleaseKey = releaseRequest.Name
	releaseHistory := new(models.ReleaseHistory)
	releaseHistory.DataChange_LastModifiedBy = releaseRequest.Operator
	releaseHistory.DataChange_CreatedBy = releaseRequest.Operator
	releaseHistory.ClusterName = releaseRequest.ClusterName
	releaseHistory.NamespaceName = releaseRequest.Name
	releaseHistory.AppId = releaseRequest.AppId
	releaseHistory.LaneName = releaseRequest.LaneName
	//查询修改的key配置并计入发布历史
	items, err := s.itemRepository.FindItemByNamespaceIdInKey(strconv.Itoa(int(releaseRequest.NamespaceId)), releaseRequest.Keys)
	if err != nil {
		return errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdInKey failed")
	}
	operationContext, err := json.Marshal(items)
	if err != nil {
		return errors.Wrap(err, "json.Marshal(items) error")
	}
	//return s.CreatePublic(release, namespaceId, keys, releaseHistory)
	//	return s.CreatePrivate(release, namespaceId, keys, releaseHistory)
	releaseMessage := new(models.ReleaseMessage)
	releaseHistory.OperationContext = string(operationContext)

	if releaseRequest.LaneName == "default" {
		releaseHistory.BranchName = "普通发布"
		releaseHistory.Operation = 0
		releaseMessage.Message = release.AppId + "+" + release.ClusterName + "+" + release.NamespaceName
	} else {
		releaseHistory.BranchName = "灰度发布"
		releaseHistory.Operation = 1
		releaseMessage.Message = release.AppId + release.LaneName + "+" + release.ClusterName + "+" + release.NamespaceName
	}

	db := s.db.Begin()
	if err := s.itemRepository.DeleteByIdOnRelease(db, strconv.Itoa(int(releaseRequest.NamespaceId)), releaseRequest.Keys); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
	}
	if err := s.itemRepository.UpdateByNamespaceId(db, strconv.Itoa(int(releaseRequest.NamespaceId)), releaseRequest.Keys); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
	}
	var itemsRelease = make([]*models.Item, 0)
	//查询的是未提交的事物内容，所以使用db在service层进行查询
	if err := db.Table(models.ItemTableName).Find(&itemsRelease, "NamespaceId=? and IsDeleted=0 and Status=1", strconv.Itoa(int(releaseRequest.NamespaceId))).Error; err != nil {
		db.Rollback()
		return errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	m := make(map[string]string)
	for i := range itemsRelease {
		m[itemsRelease[i].Key] = itemsRelease[i].ReleaseValue
	}
	config, err := json.Marshal(m)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	release.Configurations = string(config)
	releaseContext, err := json.Marshal(itemsRelease)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "json.Marshal(items) error")
	}
	releaseHistory.ReleaseContext = string(releaseContext)
	//删除以往配置
	if err := s.releaseRepository.Delete(db, release.AppId, release.ClusterName, release.NamespaceName, release.LaneName); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Delete() error")
	}
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
	}
	if err := s.releaseHistoryRepository.Create(db, releaseHistory); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseHistoryRepository.Create() error")
	}
	if err := s.repository.DeleteByMessage(db, releaseMessage.Message); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByMessage() error")
	}

	if err := s.repository.Create(db, releaseMessage); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	db.Commit()

	return nil
}

//查询该配置文件发布时的配置
func (s releaseMessageService) FindConfig(namespaceId string) (string, error) {
	m := make(map[string]string)
	items, err := s.itemRepository.FindItemByNamespaceId(namespaceId, "")
	if err != nil {
		return "", errors.Wrap(err, "call ItemRepository.Create() error")
	}
	for i := range items {
		m[items[i].Key] = items[i].Value
	}
	config, err := json.Marshal(m)
	if err != nil {
		return "", errors.Wrap(err, "call ItemRepository.Create() error")
	}
	return string(config), nil
}

//灰度全量发布，首先获取灰度配置，然后获取主版本配置，最后灰度配置覆盖主版本配置
func (s releaseMessageService) ReleaseGrayTotal(namespaceId, name, appId, cluster, operator string, isDeleted bool) error {
	items1, err := s.itemRepository.FindItemByNamespaceId(namespaceId, "") //灰度的所有配置
	if err != nil {
		return errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	//灰度信息
	appgrey, err := s.appNamespaceRepository.FindAppNamespaceById(namespaceId)
	if err != nil {
		return errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}

	//主版本信息
	app, err := s.appNamespaceRepository.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appId, cluster, "default", name)
	if err != nil {
		return errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	releaseHistory := new(models.ReleaseHistory)
	releaseHistory.DataChange_LastModifiedBy = operator
	releaseHistory.DataChange_CreatedBy = operator
	releaseHistory.ClusterName = app.ClusterName
	releaseHistory.NamespaceName = app.Name
	releaseHistory.AppId = appId
	releaseHistory.BranchName = "灰度全量发布"
	releaseHistory.Operation = 2
	releaseHistory.LaneName = app.LaneName
	release := new(models.Release)
	release.AppId = app.AppId
	release.NamespaceName = app.Name
	release.LaneName = app.LaneName
	release.ClusterName = app.ClusterName
	release.ReleaseKey = app.Name
	release.DataChange_CreatedBy = operator
	release.DataChange_LastModifiedBy = operator
	items2, err := s.itemRepository.FindItemByNamespaceId(strconv.FormatUint(app.Id, 10), "") //主版本所有配置
	items := make([]*models.Item, 0)
	m := make(map[string]int)
	for i := range items2 {
		m[items2[i].Key] = i
	}
	for i, _ := range items1 {
		if items1[i].Status != 1 {
			return errors.New("有未发布的值，只有灰度发布所有值都进行过发布，才能进行全量发布")
		}
		if j, ok := m[items1[i].Key]; ok { //item2是主版本配置，因此只需要修改发布值
			if items2[j].Value != items1[i].Value {
				items2[j].Status = 2
			}
			items2[j].Value = items1[i].Value
			items2[j].ReleaseValue = items1[i].ReleaseValue
			items2[j].Comment = items1[i].Comment
			items2[j].Describe = items1[i].Describe
			items2[j].DataChange_LastModifiedBy = operator
			items2[j].DataChange_CreatedTime = items1[i].DataChange_CreatedTime
			items = append(items, items2[j])
		} else {
			items1[i].Id = 0 //item1是灰度版本，所以需要新增操作
			items1[i].DataChange_CreatedBy = operator
			items1[i].DataChange_LastModifiedBy = operator
			items1[i].NamespaceId = app.Id
			items1[i].Status = 0
			items1[i].DataChange_CreatedTime = time.Now()
			items = append(items, items1[i])
		}
	}
	operationContext, err := json.Marshal(items)
	if err != nil {
		return errors.Wrap(err, "json.Marshal(items) error")
	}
	releaseHistory.OperationContext = string(operationContext)
	for i, _ := range items {
		items[i].Status = 1
	}

	//通知泳道
	releaseMessage := new(models.ReleaseMessage)
	releaseMessage.Message = appId + "+" + app.ClusterName + "+" + app.Name

	db := s.db.Begin()
	if err := s.itemRepository.Saves(db, items); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.Saves() error")
	}
	itemConfig := make([]*models.Item, 0)
	if err := db.Table(models.ItemTableName).Find(&itemConfig, "NamespaceId=? and IsDeleted=0 and Status=1", app.Id).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	releaseContext, err := json.Marshal(itemConfig)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "json.Marshal(items) error")
	}
	releaseHistory.ReleaseContext = string(releaseContext)

	conf := make(map[string]string)
	for _, i := range itemConfig {
		if i.ReleaseValue != "" {
			conf[i.Key] = i.ReleaseValue
		}
	}
	config, err := json.Marshal(conf)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}

	if err := s.releaseRepository.Delete(db, release.AppId, release.ClusterName, release.NamespaceName, release.LaneName); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Delete() error")
	}
	release.Configurations = string(config)
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
	}
	if isDeleted {
		if err := s.itemRepository.DeleteByNamespaceId(db, namespaceId); err != nil {
			db.Rollback()
			return errors.Wrap(err, "call itemRepository.DeleteByIdOnRelease() error")
		}
		if err := s.appNamespaceRepository.DeleteById(db, namespaceId); err != nil {
			db.Rollback()
			return errors.Wrap(err, "call AppNamespaceRepository.DeleteById() error")
		}
		if err := s.releaseRepository.Delete(db, appgrey.AppId, appgrey.ClusterName, appgrey.Name, appgrey.LaneName); err != nil {
			db.Rollback()
			return errors.Wrap(err, "call releaseRepository.Delete() error")
		}
	}
	if err := s.releaseHistoryRepository.Create(db, releaseHistory); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseHistoryRepository.Create() error")
	}
	if err := s.repository.DeleteByMessage(db, releaseMessage.Message); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByMessage() error")
	}
	if err := s.repository.Create(db, releaseMessage); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	db.Commit()
	return nil
}
