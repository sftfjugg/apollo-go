package services

import (
	"apollo-adminserivce/internal/app/adminservice/repositories"
	"apollo-adminserivce/internal/pkg/models"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

type ReleaseMessageService interface {
	Create(name, appId, clusterName, comment, namespaceName, namespaceId string, isPublic bool) error
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

func (s releaseMessageService) Create(name, appId, clusterName, comment, namespaceName, namespaceId string, isPublic bool) error {
	release := new(models.Release)
	release.NamespaceName = namespaceName
	release.AppId = appId
	release.Comment = comment
	release.ClusterName = clusterName
	release.Name = time.Now().String() + name
	release.DataChange_CreatedTime = time.Now()
	release.DataChange_LastTime = time.Now()
	configurations, err := s.FindConfig(namespaceId)
	if err != nil {
		return errors.Wrap(err, "call releaseRepository.Create() error")
	}
	release.Configurations = configurations
	if isPublic {
		return s.CreatePublic(release, namespaceId)
	} else {
		return s.CreatePrivate(release, namespaceId)
	}
	return nil
}

//私有化配置的发布，私有化配置发布只会发布对应的项目泳道（泳道在该项目中使用集群建立）
func (s releaseMessageService) CreatePrivate(release *models.Release, namespaceId string) error {
	releaseMessage := new(models.ReleaseMessage)
	releaseMessage.Message = release.AppId + "+" + release.ClusterName + "+application"
	releaseMessage.DataChange_LastTime = time.Now()
	db := s.db.Begin()
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
	}
	if err := s.itemRepository.UpdateByNamespaceId(db, namespaceId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
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

//公共配置发布，发布该项目所有泳道（泳道在该项目中使用集群建立）
func (s releaseMessageService) CreatePublic(release *models.Release, namespaceId string) error {
	appNamespaces, err := s.appNamespaceRepository.FindClusterNameByAppId(release.AppId)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceRepository.FindClusterNameByAppId() error")
	}
	//通知所有泳道的数据库模型
	releaseMessages := make([]*models.ReleaseMessage, 0)
	messaages := make([]string, 0)
	for i := range appNamespaces {
		releaseMessage := new(models.ReleaseMessage)
		releaseMessage.Message = release.AppId + "+" + appNamespaces[i].ClusterName + "+application"
		messaages = append(messaages, releaseMessage.Message)
		releaseMessage.DataChange_LastTime = time.Now()
		releaseMessages = append(releaseMessages, releaseMessage)
	}
	db := s.db.Begin()
	if err := s.releaseRepository.Create(db, release); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Create() error")
	}
	if err := s.itemRepository.UpdateByNamespaceId(db, namespaceId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.UpdateByNamespaceId() error")
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
	items, err := s.itemRepository.FindItemByNamespaceId(namespaceId)
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
