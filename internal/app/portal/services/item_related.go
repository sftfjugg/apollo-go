package services

import (
	"apollo-adminserivce/internal/app/portal/models"
	"apollo-adminserivce/internal/app/portal/repositories"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//并不关联业务层app_namespcae,关联的为
type ItemRelatedService interface {
	Create(item *models.Item) error
	Creates(items []*models.Item) error
	Update(item *models.Item) error
	DeleteById(id string) error
	DeleteByNamespaceId(namespaceId string) error
	FindItemByNamespaceId(namespaceId string) ([]*models.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error)
}

type itemRelatedService struct {
	db         *gorm.DB
	repository repositories.ItemRelatedRepisitory
}

func NewItemRelatedService(db *gorm.DB, repository repositories.ItemRelatedRepisitory) ItemRelatedService {
	return &itemRelatedService{
		db:         db,
		repository: repository,
	}
}

func (s itemRelatedService) Create(item *models.Item) error {
	items, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemRelatedService.FindOneItemByNamespaceIdAndKey() error")
	}
	if items.Key != "" {
		return errors.New("key already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Create(db, item); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRelatedRepisitory.Create() error")
	}
	db.Commit()
	return nil
}

func (s itemRelatedService) Creates(items []*models.Item) error {
	db := s.db.Begin()
	if err := s.repository.Creates(db, items); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRelatedRepisitory.Creates() error")
	}
	db.Commit()
	return nil
}

func (s itemRelatedService) Update(item *models.Item) error {
	items, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemRelatedService.FindOneItemByNamespaceIdAndKey() error")
	}
	if items.Key != "" && items.Id != item.Id {
		return errors.New("key already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Update(db, item); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRelatedRepisitory.Update() error")
	}
	db.Commit()
	return nil
}

func (s itemRelatedService) DeleteById(id string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteById(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRelatedRepisitory.DeleteById() error")
	}
	db.Commit()
	return nil
}

func (s itemRelatedService) DeleteByNamespaceId(namespaceId string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteByNamespaceId(db, namespaceId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRelatedRepisitory.DeleteByNamespaceId() error")
	}
	db.Commit()
	return nil
}

func (s itemRelatedService) FindItemByNamespaceId(namespaceId string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceId(namespaceId)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRelatedRepisitory.FindItemByNamespaceId() error")
	}
	return items, nil
}

func (s itemRelatedService) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceIdAndKey(namespaceId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRelatedRepisitory.FindItemByNamespaceIdAndKey() error")
	}
	return items, nil
}

func (s itemRelatedService) FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error) {
	item, err := s.repository.FindOneItemByNamespaceIdAndKey(namespaceId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRelatedRepisitory.FindOneItemByNamespaceIdAndKey() error")
	}
	return item, nil
}
