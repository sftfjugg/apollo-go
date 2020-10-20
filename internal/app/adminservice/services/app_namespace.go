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
	DeleteById(id string) error
	Update(appNamespace *models.AppNamespace) error
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
}

type appNamespaceService struct {
	db         *gorm.DB
	repository repositories.AppNamespaceRepository
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
	db := s.db.Begin()
	appNamespace.DataChange_LastTime = time.Now()
	appNamespace.DataChange_CreatedTime = time.Now()
	if err := s.repository.Create(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call appNamespaceService.Create() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) DeleteById(id string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteById(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call appNamespaceService.Delete() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) Update(appNamespace *models.AppNamespace) error {
	appNamespace.DataChange_LastTime = time.Now()
	db := s.db.Begin()
	if err := s.repository.Update(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call appNamespaceService.Update() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	appNamespaces, err := s.repository.FindAppNamespaceByAppIdAndClusterName(appId, clusterName)
	if err != nil {
		return nil, errors.Wrap(err, "call appNamespaceService.FindAppNamespaceByAppIdAndClusterName() error")
	}
	return appNamespaces, nil
}
