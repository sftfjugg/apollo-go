package repositories

import (
	models2 "apollo-adminserivce/internal/app/adminservice/models"
	"apollo-adminserivce/internal/pkg/models"
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

type ItemRepisitory interface {
	Create(db *gorm.DB, item *models.Item) error
	Creates(db *gorm.DB, items []*models.Item) error
	Update(db *gorm.DB, item *models.Item) error
	UpdateByNamespaceId(db *gorm.DB, namespaceId string) error
	DeleteById(db *gorm.DB, id string) error
	DeleteByNamespaceId(db *gorm.DB, namespaceId string) error
	FindItemByNamespaceId(namespaceID string) ([]*models.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error)
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
	item.DataChange_CreatedTime = time.Now()
	item.DataChange_LastTime = time.Now()
	if err := db.Create(item).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.Create failed")
	}
	return nil
}

func (r itemRepisitory) Creates(db *gorm.DB, items []*models.Item) error {
	s := "insert into Item(NamespaceId,Key,Value,Comment,Describe,DataChange_CreatedBy,DataChange_LastModifiedBy,DataChange_CreatedTime,DataChange_LastTime) values"
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(s); err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	for i, r := range items {
		if i == len(items)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s','%s');", r.NamespaceId, r.Key, r.Value, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, time.Now(), time.Now()))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s','%s'),", r.NamespaceId, r.Key, r.Value, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, time.Now(), time.Now()))
		}
	}
	if err := db.Exec(buffer.String()).Error; err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	return nil
}

func (r itemRepisitory) Update(db *gorm.DB, item *models.Item) error {
	item.DataChange_LastTime = time.Now()
	if err := db.Table(models.ItemTableName).Where("Id=?", item.Id).Update(&item).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.Update failed")
	}
	return nil
}

func (r itemRepisitory) UpdateByNamespaceId(db *gorm.DB, namespaceId string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId= ? and IsDeleted=0", namespaceId).Update("Status= 1").Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.UpdateByNamespaceId failed")
	}
	return nil
}

func (r itemRepisitory) DeleteByNamespaceId(db *gorm.DB, namespaceId string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId=? and IsDeleted=0", namespaceId).Update("IsDeleted= 1").Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteByNamespaceId failed")
	}
	return nil
}

func (r itemRepisitory) DeleteById(db *gorm.DB, id string) error {
	if err := db.Table(models.ItemTableName).Where("Id=?", id).Update("IsDeleted= 1").Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteById failed")
	}
	return nil
}

func (r itemRepisitory) FindItemByNamespaceId(namespaceID string) ([]*models.Item, error) {
	var items = make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0", namespaceID).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	return items, nil
}

func (r itemRepisitory) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	items := make([]*models.Item, 0)
	if err := r.db.Table(models.AppTableName).Find(&items, "NamespaceId =? and Key like ? and IsDeleted=0", namespaceId, "%"+key+"%").Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return items, nil
}

func (r itemRepisitory) FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error) {
	item := new(models.Item)
	if err := r.db.Table(models.AppTableName).First(&item, "NamespaceId =? and Key = ? and IsDeleted=0", namespaceId, key).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return item, nil
}

func (r itemRepisitory) FindPublicItemByAppId(appId string) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if err := r.db.Raw("select Item.key,Item.value from Item,AppNamespace where AppNamespace.IsPublic=1 and Item.NamespaceId=AppNamespace.Id and AppNamespace.AppId=? and AppNamespace.IsDeleted=0", appId).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindPublicItemByAppId failed")
	}
	return items, nil
}

func (r itemRepisitory) FindPrivateItemByAppIdandClusterName(appId, clusterName string) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if err := r.db.Raw("select Item.key,Item.value from Item,AppNamespace where AppNamespace.IsPublic=0 and Item.NamespaceId=AppNamespace.Id and AppNamespace.AppId=? and AppNamespace.ClusterName=? and AppNamespace.IsDeleted=0", appId, clusterName).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindPrivateItemByAppIdandClusterName failed")
	}
	return items, nil
}
