package services

import (
	"apollo-adminserivce/internal/app/adminservice/repositories"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

type AppNamespaceService interface {
	Create(appNamespace *models.AppNamespace) error
	CreateByRelated(appNamespace *models.AppNamespace, items []*models.Item, clusterName, appId string) error
	DeleteById(id string) error
	Update(appNamespace *models.AppNamespace) error
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
	FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name string) (*models.AppNamespace, error)
}

type appNamespaceService struct {
	db             *gorm.DB
	repository     repositories.AppNamespaceRepository
	itemRepository repositories.ItemRepisitory
	itemService    ItemService
}

func NewAppNamespaceService(
	db *gorm.DB,
	repository repositories.AppNamespaceRepository,
) AppNamespaceService {
	return &appNamespaceService{
		db:         db,
		repository: repository,
	}
}

func (s appNamespaceService) Create(appNamespace *models.AppNamespace) error {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return errors.New("name alrealy exists")
	}
	db := s.db.Begin()
	if err := s.repository.Create(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.Create() error")
	}
	db.Commit()
	return nil
}
func (s appNamespaceService) CreateByRelated(appNamespace *models.AppNamespace, items []*models.Item, clusterName, appId string) error {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return errors.New("name alrealy exists")
	}
	appNamespace.ClusterName = clusterName
	appNamespace.AppId = appId
	db := s.db.Begin()
	if err := s.repository.Create(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.Create() error")
	}
	for _, item := range items {
		item.NamespaceId = appNamespace.Id
		item.DataChange_LastTime = time.Now()
		item.DataChange_CreatedTime = time.Now()
	}
	if err := s.itemRepository.Creates(db, items); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Creates() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) DeleteById(id string) error {
	db := s.db.Begin()
	if err := s.itemRepository.DeleteById(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.DeleteById() error")
	}
	if err := s.repository.DeleteById(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.DeleteById() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) Update(appNamespace *models.AppNamespace) error {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" && app.Name != appNamespace.Name {
		return errors.New("name alrealy exists")
	}
	db := s.db.Begin()
	if err := s.repository.Update(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.Update() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error) {
	appNamespaces, err := s.repository.FindAppNamespaceByAppIdAndClusterName(appId, clusterName)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindAppNamespaceByAppIdAndClusterName() error")
	}
	return appNamespaces, nil
}

func (s appNamespaceService) FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name string) (*models.AppNamespace, error) {
	appNamespace, err := s.repository.FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	return appNamespace, nil
}
