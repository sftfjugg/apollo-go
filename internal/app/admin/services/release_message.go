package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type ReleaseMessageService interface {
	//正常发布和灰度发布
	Create(appId, clusterName, comment, name, namespaceId, operator string, isPublic bool, keys []string) error
}

type releaseMessageService struct {
	repository             repositories.ReleaseMessageRepository
	releaseRepository      repositories.Release
	appNamespaceRepository repositories.AppNamespaceRepository
	itemRepository         repositories.ItemRepisitory
	db                     *gorm.DB
}

func NewReleaseMessageService(
	repository repositories.ReleaseMessageRepository,
	releaseRepository repositories.Release,
	appNamespaceRepository repositories.AppNamespaceRepository,
	itemRepository repositories.ItemRepisitory,
	db *gorm.DB,
) ReleaseMessageService {
	return &releaseMessageService{
		repository:             repository,
		releaseRepository:      releaseRepository,
		appNamespaceRepository: appNamespaceRepository,
		itemRepository:         itemRepository,
		db:                     db,
	}
}

func (s releaseMessageService) Create(appId, clusterName, comment, name, namespaceId, operator string, isPublic bool, keys []string) error {
	release := new(models.Release)
	if appId == "" {
		app, err := s.appNamespaceRepository.FindAppNamespaceById(namespaceId)
		if err != nil {
			return errors.Wrap(err, "call appNamespaceRepository.FindAppNamespaceById() error")
		}
		appId = app.AppId
		clusterName = app.ClusterName
		isPublic = app.IsPublic
		name = app.Name
	}
	release.DataChange_LastModifiedBy = operator
	release.DataChange_CreatedBy = operator
	release.NamespaceName = name
	release.AppId = appId
	release.Comment = comment
	release.ClusterName = clusterName
	release.ReleaseKey = name
	if clusterName == "default" {
		return s.CreatePublic(release, namespaceId, keys)
	} else {
		return s.CreatePrivate(release, namespaceId, keys)
	}

	return nil
}

//泳道发布
func (s releaseMessageService) CreatePrivate(release *models.Release, namespaceId string, keys []string) error {
	releaseMessage := new(models.ReleaseMessage)
	releaseMessage.Message = release.AppId + "+" + release.ClusterName + "+" + release.NamespaceName
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
	for i := range items {
		m[items[i].Key] = items[i].ReleaseValue
	}
	config, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	release.Configurations = string(config)
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
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
func (s releaseMessageService) CreatePublic(release *models.Release, namespaceId string, keys []string) error {
	appNamespaces, err := s.appNamespaceRepository.FindClusterNameByAppId(release.AppId)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceRepository.FindClusterNameByAppId() error")
	}
	//通知所有泳道的数据库模型
	releaseMessages := make([]*models.ReleaseMessage, 0)
	messaages := make([]string, 0)
	for i := range appNamespaces {
		releaseMessage := new(models.ReleaseMessage)
		releaseMessage.Message = release.AppId + "+" + appNamespaces[i].ClusterName + "+" + release.NamespaceName
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
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	release.Configurations = string(config)
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
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
	items, err := s.itemRepository.FindItemByNamespaceId(namespaceId, "xs")
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
