package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/time"
)

//并不关联业务层app_namespcae
type ItemRelatedRepisitory interface {
	Create(db *gorm.DB, item *models.Item) error
	Creates(db *gorm.DB, item []*models.Item) error
	Update(db *gorm.DB, item *models.Item) error
	DeleteById(db *gorm.DB, id string) error
	DeleteByNamespaceId(db *gorm.DB, namespaceId string) error
	FindItemByNamespaceId(namespaceID string) ([]*models.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error)
}

type itemRelatedRepisitory struct {
	db *gorm.DB
}

func NewItemRelatedRepisitory(db *gorm.DB) ItemRelatedRepisitory {
	return &itemRelatedRepisitory{
		db: db,
	}
}

func (r itemRelatedRepisitory) Create(db *gorm.DB, item *models.Item) error {
	item.DataChange_LastTime = time.Now()
	item.DataChange_CreatedTime = time.Now()
	if err := db.Create(item).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.Create failed")
	}
	return nil
}

func (r itemRelatedRepisitory) Creates(db *gorm.DB, items []*models.Item) error {
	if err := db.Create(items).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.Creates failed")
	}
	return nil
}

func (r itemRelatedRepisitory) Update(db *gorm.DB, item *models.Item) error {
	item.DataChange_LastTime = time.Now()
	if err := db.Table(models.ItemTableName).Where("Id=?", item.Id).Update(&item).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.Update failed")
	}
	return nil
}

func (r itemRelatedRepisitory) DeleteByNamespaceId(db *gorm.DB, namespaceId string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId=? and IsDeleted=0", namespaceId).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteByNamespaceId failed")
	}
	return nil
}

func (r itemRelatedRepisitory) DeleteById(db *gorm.DB, id string) error {
	if err := db.Table(models.ItemTableName).Where("Id=?", id).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteById failed")
	}
	return nil
}

func (r itemRelatedRepisitory) FindItemByNamespaceId(namespaceID string) ([]*models.Item, error) {
	var items = make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0", namespaceID).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	return items, nil
}

func (r itemRelatedRepisitory) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	item := make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&item, "NamespaceId =? and `Key` like ? and IsDeleted=0", namespaceId, "%"+key+"%").Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return item, nil
}

func (r itemRelatedRepisitory) FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error) {
	item := new(models.Item)
	if err := r.db.Table(models.ItemTableName).First(&item, "NamespaceId =? and `Key` = ? and IsDeleted=0", namespaceId, key).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdAndKey failed")
	}
	return item, nil
}
