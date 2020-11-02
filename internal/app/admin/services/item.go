package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"sort"
)

type ItemService interface {
	Create(item *models.Item) error
	CreateOrUpdateItem(item *models.Item) error
	Creates(item []*models.Item) error
	Update(item *models.Item) error
	DeleteById(id, operator string) error
	DeleteByNamespaceId(namespaceId string) error
	FindItemByAppIdAndKey(appId, key, format, comment string) ([]*models2.AppNamespace, error)
	FindItemByNamespaceId(namespaceID, comment string) ([]*models.Item, error)
	FindItemByNamespaceIdOnRelease(namespaceID string) ([]*models.Item, error)
	FindItemByKeyForPage(key, format string, pageSize, pageNum int) (*models2.ItemPage, error)
	FindAppItemByKeyForPage(key, format string, pageSize, pageNum int) (*models2.AppNamespacePage, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error)
	FindAllComment(appId string) ([]string, error)
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
	item2, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	if item2.Key != "" {
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

func (s itemService) CreateOrUpdateItem(item *models.Item) error {
	item2, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	item.Id = item2.Id
	if item2.Key != "" {
		if err := s.Update(item); err != nil {
			return errors.Wrap(err, "call itemService.Update() error")
		}
	} else {
		item.DataChange_CreatedBy = item.DataChange_LastModifiedBy
		if err := s.Create(item); err != nil {
			return errors.Wrap(err, "call itemService.Create() error")
		}
	}
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

func (s itemService) DeleteById(id, operator string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteById(db, id, operator); err != nil {
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
	item2, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	if item2.Key != "" && item2.Key != item.Key {
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

func (s itemService) FindItemByNamespaceId(namespaceID, comment string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceId(namespaceID, comment)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	return items, nil
}

func (s itemService) FindItemByNamespaceIdOnRelease(namespaceID string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceIdOnRelease(namespaceID)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceIdOnRelease() error")
	}
	return items, nil
}

func (s itemService) FindItemByKeyForPage(key, format string, pageSize, pageNum int) (*models2.ItemPage, error) {
	items, err := s.repository.FindItemByKeyForPage(key, format, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByKeyForPage() error")
	}
	count, err := s.repository.FindItemCountByKey(key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemCountByKey() error")
	}
	itemPage := new(models2.ItemPage)
	itemPage.Items = items
	itemPage.Total = count
	return itemPage, nil
}

func (s itemService) FindAppItemByKeyForPage(key, format string, pageSize, pageNum int) (*models2.AppNamespacePage, error) {
	items, err := s.repository.FindItemByKeyForPage(key, format, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByKeyForPage() error")
	}
	appNamespaces := s.ItemChangeAppNamespace(items)
	count, err := s.repository.FindItemCountByKey(key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemCountByKey() error")
	}
	appNamespacePage := new(models2.AppNamespacePage)
	appNamespacePage.Total = count
	appNamespacePage.AppNamespaces = appNamespaces
	return appNamespacePage, nil
}

func (s itemService) FindItemByAppIdAndKey(appId, key, format, comment string) ([]*models2.AppNamespace, error) {
	items, err := s.repository.FindItemByAppIdAndKey(appId, key, format, comment)
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

func (s itemService) FindAllComment(appId string) ([]string, error) {
	items, err := s.repository.FindAllComment(appId)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindAllComment() error")
	}
	comments := make([]string, 0)
	for _, i := range items {
		if i.Comment != "" {
			comments = append(comments, i.Comment)
		}
	}
	return comments, nil
}

//作用是将Item格式转化为前端展示格式
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
				itemModel.Value = s.Value
				itemModel.Key = s.Key
				itemModel.NamespaceId = s.NamespaceId
				itemModel.DataChange_CreatedTime = s.DataChange_CreatedTime
				itemModel.DataChange_LastTime = s.DataChange_LastTime
				itemModel.DataChange_LastModifiedBy = s.DataChange_LastModifiedBy
				itemModel.DataChange_CreatedBy = s.DataChange_CreatedBy
				itemModel.Describe = s.Describe
				itemModel.Comment = s.Comment
				itemModel.ReleaseValue = s.ReleaseValue
				itemModel.Status = s.Status
				its = append(its, itemModel)
				namespace.LaneName = s.LaneName
				namespace.Id = s.NamespaceId
				appNamespace.AppId = s.AppId
				//appNamespace.
				appNamespace.Name = s.Name
				if s.Format != "" {
					appNamespace.Format = s.Format
				}
			}
			namespace.Items = its
			appNamespace.Namespaces = append(appNamespace.Namespaces, namespace)
		}
		appNamespaces = append(appNamespaces, appNamespace)
	}
	sort.Sort(models2.AppNamespaceSlice(appNamespaces))
	return appNamespaces
}
