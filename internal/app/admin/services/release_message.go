package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
	"strconv"
)

type ReleaseMessageService interface {
	//正常发布和灰度发布
	Create(appId, clusterName, comment, name, namespaceId, laneName, operator string, keys []string) error
	ReleaseGrayTotal(namespaceId, name, appId, cluster, operator string, isDeleted bool) error
}

type releaseMessageService struct {
	repository             repositories.ReleaseMessageRepository
	releaseRepository      repositories.Release
	appNamespaceRepository repositories.AppNamespaceRepository
	itemRepository         repositories.ItemRepisitory
	releaseHistoryService  repositories.ReleaseHistoryRepository
	db                     *gorm.DB
}

func NewReleaseMessageService(
	repository repositories.ReleaseMessageRepository,
	releaseRepository repositories.Release,
	appNamespaceRepository repositories.AppNamespaceRepository,
	releaseHistoryService repositories.ReleaseHistoryRepository,
	itemRepository repositories.ItemRepisitory,
	db *gorm.DB,
) ReleaseMessageService {
	return &releaseMessageService{
		repository:             repository,
		releaseRepository:      releaseRepository,
		appNamespaceRepository: appNamespaceRepository,
		itemRepository:         itemRepository,
		releaseHistoryService:  releaseHistoryService,
		db:                     db,
	}
}

func (s releaseMessageService) Create(appId, clusterName, comment, name, namespaceId, laneName, operator string, keys []string) error {
	release := new(models.Release)
	if appId == "" {
		app, err := s.appNamespaceRepository.FindAppNamespaceById(namespaceId)
		if err != nil {
			return errors.Wrap(err, "call appNamespaceRepository.FindAppNamespaceById() error")
		}
		appId = app.AppId
		clusterName = app.ClusterName
		laneName = app.LaneName
		name = app.Name
	}
	release.DataChange_LastModifiedBy = operator
	release.DataChange_CreatedBy = operator
	release.NamespaceName = name
	release.AppId = appId
	release.Comment = comment
	release.LaneName = laneName
	release.ClusterName = clusterName
	release.ReleaseKey = name
	releaseHistory := new(models.ReleaseHistory)
	releaseHistory.DataChange_LastModifiedBy = operator
	releaseHistory.DataChange_CreatedBy = operator
	releaseHistory.ClusterName = clusterName
	releaseHistory.NamespaceName = name
	releaseHistory.AppId = appId
	releaseHistory.LaneName = laneName
	//查询修改的key配置并计入发布历史
	items, err := s.itemRepository.FindItemByNamespaceIdInKey(namespaceId, keys)
	if err != nil {
		return errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdInKey failed")
	}
	operationContext, err := json.Marshal(items)
	if err != nil {
		return errors.Wrap(err, "json.Marshal(items) error")
	}
	releaseHistory.OperationContext = string(operationContext)

	if laneName == "default" {
		return s.CreatePublic(release, namespaceId, keys, releaseHistory)
	} else {
		return s.CreatePrivate(release, namespaceId, keys, releaseHistory)
	}

	return nil
}

//泳道发布
func (s releaseMessageService) CreatePrivate(release *models.Release, namespaceId string, keys []string, releaseHistory *models.ReleaseHistory) error {
	releaseMessage := new(models.ReleaseMessage)
	releaseMessage.Message = release.AppId + release.LaneName + "+" + release.ClusterName + "+" + release.NamespaceName
	db := s.db.Begin()
	if err := s.itemRepository.DeleteByIdOnRelease(db, namespaceId, keys); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
	}
	if err := s.itemRepository.UpdateByNamespaceId(db, namespaceId, keys); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
	}
	var items = make([]*models.Item, 0)
	//查询的是未提交的事物内容，所以使用db在service层进行查询
	if err := db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0 and Status=1", namespaceId).Error; err != nil {
		db.Rollback()
		return errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	m := make(map[string]string)
	for i := range items {
		m[items[i].Key] = items[i].ReleaseValue
	}
	config, err := json.Marshal(m)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	release.Configurations = string(config)
	releaseContext, err := json.Marshal(items)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "json.Marshal(items) error")
	}
	releaseHistory.ReleaseContext = string(releaseContext)
	releaseHistory.BranchName = "灰度发布"
	releaseHistory.Operation = 1
	//删除以往配置
	if err := s.releaseRepository.Delete(db, release.AppId, release.ClusterName, release.NamespaceName, release.LaneName); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Delete() error")
	}
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
	}
	if err := s.releaseHistoryService.Create(db, releaseHistory); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseHistoryService.Create() error")
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

//正常发布，正常发布分为正常发布和通过灰度进行全量发布
func (s releaseMessageService) CreatePublic(release *models.Release, namespaceId string, keys []string, releaseHistory *models.ReleaseHistory) error {
	appNamespaces, err := s.appNamespaceRepository.FindClusterNameByAppId(release.AppId)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceRepository.FindClusterNameByAppId() error")
	}
	//通知所有泳道的数据库模型
	releaseMessages := make([]*models.ReleaseMessage, 0)
	messaages := make([]string, 0)
	for i := range appNamespaces {
		releaseMessage := new(models.ReleaseMessage)
		appId := appNamespaces[i].AppId
		if appNamespaces[i].LaneName != "default" {
			appId = appNamespaces[i].AppId + appNamespaces[i].LaneName
		}
		releaseMessage.Message = appId + "+" + appNamespaces[i].ClusterName + "+" + release.NamespaceName
		messaages = append(messaages, releaseMessage.Message)
		releaseMessages = append(releaseMessages, releaseMessage)
	}
	db := s.db.Begin()
	if err := s.itemRepository.DeleteByIdOnRelease(db, namespaceId, keys); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
	}
	if err := s.itemRepository.UpdateByNamespaceId(db, namespaceId, keys); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
	}
	var items = make([]*models.Item, 0)
	if err := db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0 and Status=1", namespaceId).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	m := make(map[string]string)
	for _, i := range items {
		if i.ReleaseValue != "" {
			m[i.Key] = i.ReleaseValue
		}
	}
	config, err := json.Marshal(m)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	release.Configurations = string(config)
	releaseHistory.BranchName = "普通发布"
	releaseHistory.Operation = 0
	releaseContext, err := json.Marshal(items)
	if err != nil {
		db.Rollback()
		return errors.Wrap(err, "json.Marshal(items) error")
	}
	releaseHistory.ReleaseContext = string(releaseContext)
	//先删除以前发布
	if err := s.releaseRepository.Delete(db, release.AppId, release.ClusterName, release.NamespaceName, release.LaneName); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Delete() error")
	}
	//发布，保证只有一个发布
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
	}
	if err := s.releaseHistoryService.Create(db, releaseHistory); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseHistoryService.Create() error")
	}
	if err := s.repository.DeleteByMessages(db, messaages); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByMessage() error")
	}
	if err := s.repository.Creates(db, releaseMessages); err != nil {
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
	appNamespaces, err := s.appNamespaceRepository.FindClusterNameByAppId(release.AppId)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceRepository.FindClusterNameByAppId() error")
	}
	//通知所有泳道的数据库模型
	releaseMessages := make([]*models.ReleaseMessage, 0)
	messaages := make([]string, 0)
	for i := range appNamespaces {
		releaseMessage := new(models.ReleaseMessage)
		appId := appNamespaces[i].AppId
		if appNamespaces[i].LaneName != "default" {
			appId = appNamespaces[i].AppId + appNamespaces[i].LaneName
		}
		releaseMessage.Message = appId + "+" + appNamespaces[i].ClusterName + "+" + release.NamespaceName
		messaages = append(messaages, releaseMessage.Message)
		releaseMessages = append(releaseMessages, releaseMessage)
	}

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
	if err := s.releaseHistoryService.Create(db, releaseHistory); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseHistoryService.Create() error")
	}
	if err := s.repository.DeleteByMessages(db, messaages); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByMessage() error")
	}
	if err := s.repository.Creates(db, releaseMessages); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	db.Commit()
	return nil
}
