package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/repositories"
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
	db                    *gorm.DB
	repository            repositories.AppNamespaceRelatedRepository
	itemRelatedRepository repositories.ItemRelatedRepisitory
}

func NewAppNamespaceRelatedService(
	db *gorm.DB,
	repository repositories.AppNamespaceRelatedRepository,
	itemRelatedRepository repositories.ItemRelatedRepisitory,
) AppNamespaceRelatedService {
	return &appNamespaceRelatedService{
		db:                    db,
		repository:            repository,
		itemRelatedRepository: itemRelatedRepository,
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
	if err := s.repository.Create(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRelatedRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceRelatedService) Delete(id string) error {
	db := s.db.Begin()
	if err := s.itemRelatedRepository.DeleteByNamespaceId(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRelatedRepository.DeleteByNamespaceId() error")
	}
	if err := s.repository.Delete(db, id); err != nil {
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
	if app.Id != appNamespace.Id && app.Name != "" {
		return errors.New("appNamespace.Name already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Update(db, appNamespace); err != nil {
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
