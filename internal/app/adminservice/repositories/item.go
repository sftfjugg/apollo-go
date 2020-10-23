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
	UpdateByNamespaceId(db *gorm.DB, namespaceId string, keys []string) error
	DeleteById(db *gorm.DB, id string) error
	DeleteByNamespaceId(db *gorm.DB, namespaceId string) error
	FindItemByNamespaceId(namespaceID string) ([]*models.Item, error)
	FindItemByKeyForPage(key string, pageSize, pageNum int) ([]*models2.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindItemByAppIdAndKey(appId, key string) ([]*models2.Item, error)
	FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error)
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
	if err := db.Create(&item).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.Create failed")
	}
	return nil
}

func (r itemRepisitory) Creates(db *gorm.DB, items []*models.Item) error {
	s := "insert into Item(`NamespaceId`,`Key`,`Value`,`Comment`,`Describe`,`DataChange_CreatedBy`,`DataChange_LastModifiedBy`,`DataChange_CreatedTime`,`DataChange_LastTime`) values"
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(s); err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	for i, r := range items {
		if i == len(items)-1 {
			buffer.WriteString(fmt.Sprintf("('%v','%s','%s','%s','%s','%s','%s','%s','%s');", r.NamespaceId, r.Key, r.Value, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, time.Now(), time.Now()))
		} else {
			buffer.WriteString(fmt.Sprintf("('%v','%s','%s','%s','%s','%s','%s','%s','%s'),", r.NamespaceId, r.Key, r.Value, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, time.Now(), time.Now()))
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

func (r itemRepisitory) UpdateByNamespaceId(db *gorm.DB, namespaceId string, keys []string) error {
	key := "('" + keys[0] + "'"
	for i := 1; i < len(keys); i++ {
		key += ",'" + keys[i] + "'"
	}
	key += ")"
	if err := db.Table(models.ItemTableName).Where("NamespaceId= ? and IsDeleted=0 and `Key` in "+key, namespaceId).Update("Status", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.UpdateByNamespaceId failed")
	}
	return nil
}

func (r itemRepisitory) DeleteByNamespaceId(db *gorm.DB, namespaceId string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId=? and IsDeleted=0", namespaceId).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteByNamespaceId failed")
	}
	return nil
}

func (r itemRepisitory) DeleteById(db *gorm.DB, id string) error {
	if err := db.Table(models.ItemTableName).Where("Id=?", id).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteById failed")
	}
	return nil
}

func (r itemRepisitory) FindItemByNamespaceId(namespaceID string) ([]*models.Item, error) {
	var items = make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0 and Status=1", namespaceID).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	return items, nil
}

func (r itemRepisitory) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	items := make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId =? and `Key` like ? and IsDeleted=0", namespaceId, "%"+key+"%").Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return items, nil
}

func (r itemRepisitory) FindItemByAppIdAndKey(appId, key string) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if err := r.db.Raw("Select I.Id,I.Key,I.Value,I.NamespaceId,A.Name,A.AppId,A.AppName,A.ClusterName,A.LaneName,I.Status,I.Comment,I.Describe,I.DataChange_CreatedBy,I.DataChange_LastModifiedBy,I.DataChange_CreatedTime,I.DataChange_LastTime from `AppNamespace` A,`Item` I where I.Key like ? and A.Id=I.NamespaceId and A.AppId=? and I.IsDeleted=0;", "%"+key+"%", appId).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return items, nil
}

func (r itemRepisitory) FindItemByKeyForPage(key string, pageSize, pageNum int) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if err := r.db.Raw("Select I.Id,I.Key,I.Value,I.NamespaceId,A.Name,A.AppId,A.AppName,A.ClusterName,A.LaneName,I.Status,I.Comment,I.Describe,I.DataChange_CreatedBy,I.DataChange_LastModifiedBy,I.DataChange_CreatedTime,I.DataChange_LastTime from `AppNamespace` A,`Item` I where I.Key like ? and A.Id=I.NamespaceId and I.IsDeleted=0 Limit ?,?;", "%"+key+"%", pageSize*(pageNum-1), pageSize).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return items, nil
}

func (r itemRepisitory) FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error) {
	item := new(models.Item)
	if err := r.db.Table(models.ItemTableName).First(&item, "NamespaceId =? and `Key` = ? and IsDeleted=0", namespaceId, key).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return item, nil
}
