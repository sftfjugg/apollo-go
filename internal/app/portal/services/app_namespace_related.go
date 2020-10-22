package services

import (
	"apollo-adminserivce/internal/app/portal/models"
	"apollo-adminserivce/internal/app/portal/repositories"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type AppNamespaceRelatedService interface {
	Create(appNamespace *models.AppNamespace) error
	Delete(id string) error
	Update(appNamespace *models.AppNamespace) error
	FindAppNamespaceByName(name string) (*models.AppNamespace, error)
	FindAppNamespaceByNameForPage(name string, pageSize, pageNum int) ([]*models.AppNamespace, error)
	FindAppNamespaceByDepartmentForPage(department string, pageSize, pageNum int) ([]*models.AppNamespace, error)
}

type appNamespaceRelatedService struct {
	db         *gorm.DB
	repository repositories.AppNamespaceRelatedRepository
}

func NewAppNamespaceRelatedService(
	db *gorm.DB,
	repository repositories.AppNamespaceRelatedRepository,
) AppNamespaceRelatedService {
	return &appNamespaceRelatedService{
		db:         db,
		repository: repository,
	}
}

func (s appNamespaceRelatedService) Create(appNamespace *models.AppNamespace) error {
	app, err := s.FindAppNamespaceByName(appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceRelatedService.FindAppNamespaceByName() error")
	}
	if app.Name != "" {
		return errors.New("appNamespace.Name already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Create(s.db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRelatedRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceRelatedService) Delete(id string) error {
	db := s.db.Begin()
	if err := s.repository.Delete(s.db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRelatedRepository.Delete() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceRelatedService) Update(appNamespace *models.AppNamespace) error {
	app, err := s.FindAppNamespaceByName(appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceRelatedService.FindAppNamespaceByName() error")
	}
	if app.Name != appNamespace.Name && app.Name != "" {
		return errors.New("appNamespace.Name already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Update(s.db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRelatedRepository.Update() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceRelatedService) FindAppNamespaceByNameForPage(name string, pageSize, pageNum int) ([]*models.AppNamespace, error) {
	appNamespaces, err := s.repository.FindAppNamespaceByNameForPage(name, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRelatedRepository.FindAppNamespaceByNameForPage() error")
	}
	return appNamespaces, nil
}

func (s appNamespaceRelatedService) FindAppNamespaceByName(name string) (*models.AppNamespace, error) {
	appNamespace, err := s.repository.FindAppNamespaceByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRelatedRepository.FindAppNamespaceByName() error")
	}
	return appNamespace, nil
}

func (s appNamespaceRelatedService) FindAppNamespaceByDepartmentForPage(department string, pageSize, pageNum int) ([]*models.AppNamespace, error) {
	appNamespaces, err := s.repository.FindAppNamespaceByDepartmentForPage(department, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRelatedRepository.FindAppNamespaceByDepartmentForPage() error")
	}
	return appNamespaces, nil
}
