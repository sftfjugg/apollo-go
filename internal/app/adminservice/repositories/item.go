package repositories

import (
	models2 "apollo-adminserivce/internal/app/adminservice/models"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ItemRepisitory interface {
	Create(db *gorm.DB, item *models.Item) error
	Creates(db *gorm.DB, items []*models.Item) error
	Update(db *gorm.DB, item *models.Item) error
	UpdateByNamespaceId(db *gorm.DB, namespaceId string) error
	DeleteByNamespaceIdAndKey(db *gorm.DB, namespaceId, key string) error
	DeleteByNamespaceId(db *gorm.DB, namespaceId string) error
	FindItemByNamespaceId(namespaceID string) ([]*models.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindPublicItemByAppId(appId string) ([]*models2.Item, error)
	FindPrivateItemByAppIdandClusterName(appId, clusterName string) ([]*models2.Item, error)
}

type itemRepisitory struct {
	db *gorm.DB
}

func NewItemRepisitory(db *gorm.DB) ItemRepisitory {
	return &itemRepisitory{
		db: db,
	}
}

func (r itemRepisitory) Create(db *gorm.DB, item *models.Item) error {
	if err := db.Create(item).Error; err != nil {
		return errors.Wrap(err, "create item error")
	}
	return nil
}

func (r itemRepisitory) Creates(db *gorm.DB, items []*models.Item) error {
	if err := db.Create(items).Error; err != nil {
		return errors.Wrap(err, "create item error")
	}
	return nil
}

func (r itemRepisitory) Update(db *gorm.DB, item *models.Item) error {
	if err := db.Table(models.ItemTableName).Where("Id=?", item.Id).Update(&item).Error; err != nil {
		return errors.Wrap(err, "update item error")
	}
	return nil
}

func (r itemRepisitory) UpdateByNamespaceId(db *gorm.DB, namespaceId string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId= ?", namespaceId).Update("Status= 1").Error; err != nil {
		return errors.Wrap(err, "update item by NamespaceId  error")
	}
	return nil
}

func (r itemRepisitory) DeleteByNamespaceId(db *gorm.DB, namespaceId string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId=?", namespaceId).Update("IsDeleted= ?", true).Error; err != nil {
		return errors.Wrap(err, "delete item by NamespaceId and Key error")
	}
	return nil
}

func (r itemRepisitory) DeleteByNamespaceIdAndKey(db *gorm.DB, namespaceId, key string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId=? and Key=?", namespaceId, key).Update("IsDeleted= ?", true).Error; err != nil {
		return errors.Wrap(err, "delete item by NamespaceId and Key error")
	}
	return nil
}

func (r itemRepisitory) FindItemByNamespaceId(namespaceID string) ([]*models.Item, error) {
	var items = make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0", namespaceID).Error; err != nil {
		return nil, errors.Wrap(err, "find item by NamespaceId error")
	}
	return items, nil
}

func (r itemRepisitory) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	item := make([]*models.Item, 0)
	if err := r.db.Table(models.AppTableName).Find(&item, "NamespaceId =? and Key like ? and IsDeleted=0", namespaceId, key).Error; err != nil {
		return nil, errors.Wrap(err, "find item by NamespaceId and Key error")
	}
	return item, nil
}

func (r itemRepisitory) FindPublicItemByAppId(appId string) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if err := r.db.Raw("select Item.key,Item.value from Item,AppNamespace where AppNamespace.IsPublic=1 and Item.NamespaceId=AppNamespace.Id and AppNamespace.AppId=? and AppNamespace.IsDeleted=0", appId).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "find item by NamespaceId and Key error")
	}
	return items, nil
}

func (r itemRepisitory) FindPrivateItemByAppIdandClusterName(appId, clusterName string) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if err := r.db.Raw("select Item.key,Item.value from Item,AppNamespace where AppNamespace.IsPublic=0 and Item.NamespaceId=AppNamespace.Id and AppNamespace.AppId=? and AppNamespace.ClusterName=? and AppNamespace.IsDeleted=0", appId, clusterName).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "find item by NamespaceId and Key error")
	}
	return items, nil
}
