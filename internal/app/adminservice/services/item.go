package services

import (
	models2 "apollo-adminserivce/internal/app/adminservice/models"
	"apollo-adminserivce/internal/app/adminservice/repositories"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ItemService interface {
	Create(item *models.Item) error
	Creates(item []*models.Item) error
	Update(item *models.Item) error
	DeleteById(id string) error
	DeleteByNamespaceId(namespaceId string) error
	FindItemByAppIdAndKey(appId, key string) ([]*models2.AppNamespace, error)
	FindItemByNamespaceId(namespaceID string) ([]*models.Item, error)
	FindItemByKeyForPage(key string, pageSize, pageNum int) ([]*models2.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error)
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
	items, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	if items.Key != "" {
		return errors.New("item already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Create(db, item); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s itemService) Creates(items []*models.Item) error {
	db := s.db.Begin()
	if err := s.repository.Creates(db, items); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Creates() error")
	}
	db.Commit()
	return nil
}

func (s itemService) DeleteById(id string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteById(db, id); err != nil {
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
	items, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	if items.Key != "" && items.Key != item.Key {
		return errors.New("item already exists")
	}
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

func (s itemService) FindItemByKeyForPage(key string, pageSize, pageNum int) ([]*models2.Item, error) {
	items, err := s.repository.FindItemByKeyForPage(key, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByKeyForPage() error")
	}
	return items, nil
}

func (s itemService) FindItemByAppIdAndKey(appId, key string) ([]*models2.AppNamespace, error) {
	items, err := s.repository.FindItemByAppIdAndKey(appId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByAppIdAndKey() error")
	}
	appNamespaces := s.ItemChangeAppNamespace(items)
	return appNamespaces, nil
}

func (s itemService) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceIdAndKey(namespaceId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	return items, nil
}

func (s itemService) FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error) {
	item, err := s.repository.FindOneItemByNamespaceIdAndKey(namespaceId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindOneItemByNamespaceIdAndKey() error")
	}
	return item, nil
}

func (s itemService) ItemChangeAppNamespace(items []*models2.Item) []*models2.AppNamespace {
	names := make(map[string][]int)
	for i, n := range items {
		names[n.Name] = append(names[n.Name], i)
	}
	appNamespaces := make([]*models2.AppNamespace, 0)
	for _, v := range names {
		clusters := make(map[string][]*models2.Item)
		for i := range v {
			clusters[items[i].ClusterName] = append(clusters[items[i].ClusterName], items[i])
		}
		appNamespace := new(models2.AppNamespace)
		for key, c := range clusters {
			namespace := new(models2.Namespace)
			namespace.ClusterName = key
			its := make([]*models.Item, 0)
			for _, s := range c {
				itemModel := new(models.Item)
				itemModel.Id = s.Id
				itemModel.Key = s.Key
				itemModel.NamespaceId = s.NamespaceId
				itemModel.DataChange_CreatedTime = s.DataChange_CreatedTime
				itemModel.DataChange_LastTime = s.DataChange_LastTime
				itemModel.DataChange_LastModifiedBy = s.DataChange_LastModifiedBy
				itemModel.DataChange_CreatedBy = s.DataChange_CreatedBy
				itemModel.Describe = s.Describe
				itemModel.Comment = s.Comment
				itemModel.Status = s.Status
				its = append(its, itemModel)
				namespace.LaneName = s.LaneName
				appNamespace.AppId = s.AppId
				appNamespace.AppName = s.AppName
				appNamespace.Name = s.Name
			}
			namespace.Items = its
			appNamespace.Namespaces = append(appNamespace.Namespaces, namespace)
		}
		appNamespaces = append(appNamespaces, appNamespace)
	}
	return appNamespaces
}
