package services

import (
	"apollo-adminserivce/internal/app/adminservice/repositories"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

type ItemService interface {
	Create(item *models.Item) error
	Update(item *models.Item) error
	DeleteByNamespaceIdAndKey(namespaceId, key string) error
	DeleteByNamespaceId(namespaceId string) error
	FindItemByNamespaceId(namespaceID string) ([]*models.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
}

type itemService struct {
	repository repositories.ItemRepisitory
	db         *gorm.DB
}

func NewItemService(
	repository repositories.ItemRepisitory,
	db *gorm.DB,
) ItemService {
	return &itemService{
		db:         db,
		repository: repository,
	}
}

func (s itemService) Create(item *models.Item) error {
	item.DataChange_CreatedTime = time.Now()
	item.DataChange_LastTime = time.Now()
	db := s.db.Begin()
	if err := s.repository.Create(db, item); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s itemService) DeleteByNamespaceIdAndKey(namespaceId, key string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteByNamespaceIdAndKey(db, namespaceId, key); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByNamespaceIdAndKey() error")
	}
	db.Commit()
	return nil
}

func (s itemService) DeleteByNamespaceId(namespaceId string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteByNamespaceId(db, namespaceId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByNamespaceIdAndKey() error")
	}
	db.Commit()
	return nil
}

func (s itemService) Update(item *models.Item) error {
	item.DataChange_LastTime = time.Now()
	db := s.db.Begin()
	if err := s.repository.Update(db, item); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Update() error")
	}
	db.Commit()
	return nil
}

func (s itemService) FindItemByNamespaceId(namespaceID string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceId(namespaceID)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	return items, nil
}

func (s itemService) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceIdAndKey(namespaceId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	return items, nil
}
