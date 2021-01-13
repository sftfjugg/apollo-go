package repositories

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type ItemRepisitory interface {
	Create(db *gorm.DB, item *models.Item) error
	Creates(db *gorm.DB, items []*models.Item) error
	Update(db *gorm.DB, item *models.Item) error
	Saves(db *gorm.DB, items []*models.Item) error
	UpdateByNamespaceId(db *gorm.DB, namespaceId string, keys []string) error
	DeleteById(db *gorm.DB, id, operator string) error
	DeleteByIdOnRelease(db *gorm.DB, namespaceId string, keys []string) error
	DeleteByNamespaceId(db *gorm.DB, namespaceId string) error
	DeleteByNamespaceIds(db *gorm.DB, namespaceIds []string) error
	FindItemByNamespaceId(namespaceID, comment string) ([]*models.Item, error)
	FindItemByNamespaceIdOnRelease(namespaceID string) ([]*models.Item, error)
	FindItemByKeyForPage(cluster, key, format string, pageSize, pageNum int) ([]*models2.Item, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindItemByNamespaceIdInKey(namespaceId string, keys []string) ([]*models.Item, error)
	FindItemByAppIdAndKey(appId, cluster, key, format, comment string) ([]*models2.Item, error)
	FindItemCountByKey(key string) (int, error)
	FindAllComment(appId string) ([]*models.Item, error) //查询所有不同标签
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
	item.Status = 0
	item.DataChange_CreatedTime = time.Now()
	item.DataChange_LastTime = time.Now()
	if err := db.Create(&item).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.Create failed")
	}
	return nil
}

func (r itemRepisitory) Creates(db *gorm.DB, items []*models.Item) error {
	s := "insert into Item(`NamespaceId`,`Key`,`Value`,`ReleaseValue`,`Status`,`Comment`,`Describe`,`DataChange_CreatedBy`,`DataChange_LastModifiedBy`,`DataChange_CreatedTime`,`DataChange_LastTime`) values"
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(s); err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	for i, r := range items {
		r.Status = 0
		if i == len(items)-1 {
			buffer.WriteString(fmt.Sprintf("('%v','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s');", r.NamespaceId, r.Key, r.Value, r.ReleaseValue, r.Status, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, time.Now(), time.Now()))
		} else {
			buffer.WriteString(fmt.Sprintf("('%v','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s'),", r.NamespaceId, r.Key, r.Value, r.ReleaseValue, r.Status, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, time.Now(), time.Now()))
		}
	}
	if err := db.Exec(buffer.String()).Error; err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	return nil
}

func (r itemRepisitory) Update(db *gorm.DB, item *models.Item) error {
	item.DataChange_LastTime = time.Now()
	item.Status = 2
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
	if err := db.Table(models.ItemTableName).Where("NamespaceId= ? and IsDeleted=0 and `Key` in "+key, namespaceId).Update("ReleaseValue`=Value,`Status", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.UpdateByNamespaceId failed")
	}
	return nil
}

func (r itemRepisitory) Saves(db *gorm.DB, items []*models.Item) error {
	s := "replace into Item(`Id`,`NamespaceId`,`Key`,`Value`,`ReleaseValue`,`Status`,`Comment`,`Describe`,`DataChange_CreatedBy`,`DataChange_LastModifiedBy`,`DataChange_CreatedTime`,`DataChange_LastTime`) values"
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(s); err != nil {
		return errors.Wrap(err, "creates items error")
	}
	for i, r := range items {
		if i == len(items)-1 {
			buffer.WriteString(fmt.Sprintf("('%v','%v','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s');", r.Id, r.NamespaceId, r.Key, r.Value, r.ReleaseValue, r.Status, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, r.DataChange_CreatedTime, time.Now()))
		} else {
			buffer.WriteString(fmt.Sprintf("('%v','%v','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s'),", r.Id, r.NamespaceId, r.Key, r.Value, r.ReleaseValue, r.Status, r.Comment, r.Describe, r.DataChange_CreatedBy, r.DataChange_LastModifiedBy, r.DataChange_CreatedTime, time.Now()))
		}
	}
	if err := db.Exec(buffer.String()).Error; err != nil {
		return errors.Wrap(err, "save item error")
	}
	return nil
}

func (r itemRepisitory) DeleteByNamespaceId(db *gorm.DB, namespaceId string) error {
	if err := db.Table(models.ItemTableName).Where("NamespaceId=? and IsDeleted=0", namespaceId).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteByNamespaceId failed")
	}
	return nil
}

func (r itemRepisitory) DeleteByNamespaceIds(db *gorm.DB, namespaceIds []string) error {
	ids := "('" + namespaceIds[0] + "'"
	for i := 1; i < len(namespaceIds); i++ {
		ids += ",'" + namespaceIds[i] + "'"
	}
	ids += ")"
	if err := db.Table(models.ItemTableName).Where("IsDeleted=0 and NamespaceId in"+ids).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteByNamespaceId failed")
	}
	return nil
}

//只是修改标记为删除
func (r itemRepisitory) DeleteById(db *gorm.DB, id, operator string) error {
	if err := db.Table(models.ItemTableName).Where("Id=?", id).Update("Status", 3, "DataChange_LastModifiedBy", operator).Error; err != nil {
		return errors.Wrap(err, "ItemRepisitory.DeleteById failed")
	}
	return nil
}

//真实删除，发布和删除配置文件时使用真实删除
func (r itemRepisitory) DeleteByIdOnRelease(db *gorm.DB, id string, keys []string) error {
	key := "('" + keys[0] + "'"
	for i := 1; i < len(keys); i++ {
		key += ",'" + keys[i] + "'"
	}
	key += ")"
	if err := db.Table(models.ItemTableName).Where("namespaceId=? and Status=3 and `Key` in "+key, id).Update("IsDeleted", 1).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return errors.Wrap(err, "ItemRepisitory.DeleteByIdOnRelease failed")
	}
	return nil
}

func (r itemRepisitory) FindItemByNamespaceId(namespaceID, comment string) ([]*models.Item, error) {
	var items = make([]*models.Item, 0)
	if comment != "" {
		comment = "and Comment='" + comment + "'"
	}
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0 "+comment+"", namespaceID).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	return items, nil
}

//获得发布时需要发布的key
func (r itemRepisitory) FindItemByNamespaceIdOnRelease(namespaceID string) ([]*models.Item, error) {
	var items = make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId=? and IsDeleted=0 and Status!=1", namespaceID).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdOnRelease failed")
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

func (r itemRepisitory) FindItemByNamespaceIdInKey(namespaceId string, keys []string) ([]*models.Item, error) {
	items := make([]*models.Item, 0)
	if err := r.db.Table(models.ItemTableName).Find(&items, "NamespaceId =? and `Key` in (?) and IsDeleted=0", namespaceId, keys).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceIdInKey failed")
	}
	return items, nil
}

func (r itemRepisitory) FindItemByAppIdAndKey(appId, cluster, key, format, comment string) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if format != "" {
		format = "and A.Format='" + format + "'  "
	}
	if comment != "" {
		comment = "and I.Comment='" + comment + "'  "
	}
	if cluster != "" {
		cluster = "and A.ClusterName='" + cluster + "'  "
	}
	if err := r.db.Raw("Select I.Id,I.Key,I.Value,I.NamespaceId,A.Name,A.AppId,A.AppName,A.ClusterName,A.LaneName,A.DeptName,A.IsDisplay,A.IsPublic,A.Format,I.Status,I.Comment,A.Comment NamespaceComment,I.Describe,I.DataChange_CreatedBy,I.DataChange_LastModifiedBy,I.DataChange_CreatedTime,I.DataChange_LastTime from `AppNamespace` A,`Item` I where I.Key like ? and A.Id=I.NamespaceId and I.IsDeleted=0 and AppId=? "+format+comment+cluster+"", "%"+key+"%", appId).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByNamespaceId failed")
	}
	return items, nil
}

func (r itemRepisitory) FindItemByKeyForPage(cluster, key, format string, pageSize, pageNum int) ([]*models2.Item, error) {
	items := make([]*models2.Item, 0)
	if format != "" {
		format = "and Format='" + format + "'  "
	}
	if cluster != "" {
		cluster = "and A.ClusterName='" + cluster + "'  "
	}
	if err := r.db.Raw("Select I.Id,I.Key,I.Value,I.NamespaceId,A.Name,A.AppId,A.AppName,A.ClusterName,A.LaneName,A.IsPublic,A.Format,I.Status,I.Comment,I.Describe,I.DataChange_CreatedBy,I.DataChange_LastModifiedBy,I.DataChange_CreatedTime,I.DataChange_LastTime,A.DeptName,A.IsDisplay from `AppNamespace` A,`Item` I where I.Key like ? and A.Id=I.NamespaceId and I.IsDeleted=0 "+format+cluster+" order by I.NamespaceId Limit ?,?;", "%"+key+"%", pageSize*(pageNum-1), pageSize).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindItemByKeyForPage failed")
	}
	return items, nil
}

func (r itemRepisitory) FindItemCountByKey(key string) (int, error) {

	var count = new(models2.Count)
	if err := r.db.Raw("Select count(*) as count  from `AppNamespace` A,`Item` I where I.Key like ? and A.Id=I.NamespaceId and I.IsDeleted=0;", "%"+key+"%").Scan(&count).Error; err != nil {
		return 0, errors.Wrap(err, "ItemRepisitory.FindItemCountByKey failed")
	}
	return count.Count, nil
}

func (r itemRepisitory) FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error) {
	item := new(models.Item)
	if err := r.db.Table(models.ItemTableName).First(&item, "NamespaceId =? and `Key` = ? and IsDeleted=0", namespaceId, key).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "ItemRepisitory.FindOneItemByNamespaceIdAndKey failed")
	}
	return item, nil
}

func (r itemRepisitory) FindAllComment(appId string) ([]*models.Item, error) {
	items := make([]*models.Item, 0)
	if err := r.db.Raw("Select I.Comment from `AppNamespace` A,`Item` I where  A.Id=I.NamespaceId and A.AppId=? and I.IsDeleted=0 group by I.Comment;", appId).Scan(&items).Error; err != nil {
		return nil, errors.Wrap(err, "ItemRepisitory.FindAllComment failed")
	}
	return items, nil
}
