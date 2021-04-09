package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type DingdingRepository interface {
	Create(dingding *models.Dingding) error
	FindAll(pageNum, pageSize int) ([]*models.Dingding, error)
	FindCount() (int, error)
	Update(dingding *models.Dingding) error
	Delete(id int) error
	Find(Type, deptName, env string, level int) (*models.Dingding, error)
	FindByName(name string) bool
}

type dingdingRepository struct {
	db *gorm.DB
}

func NewDingdingRepository(db *gorm.DB) DingdingRepository {
	return &dingdingRepository{db: db}
}

func (d dingdingRepository) Create(dingding *models.Dingding) error {
	dingding.DataChange_CreatedTime = time.Now()
	db := d.db.Begin()
	if err := db.Table(models.DingdingTableName).Create(&dingding).Error; err != nil {
		db.Rollback()
		return errors.Wrap(err, "create dingding error")
	}
	db.Commit()
	return nil
}

func (d dingdingRepository) FindAll(pageNum, pageSize int) ([]*models.Dingding, error) {
	dingdings := make([]*models.Dingding, 0)
	if err := d.db.Table(models.DingdingTableName).Limit(pageSize).Offset(pageSize*(pageNum-1)).Find(&dingdings, "IsDeleted=0").Error; err != nil {
		return nil, errors.Wrap(err, "find all dingding error")
	}
	return dingdings, nil
}

func (d dingdingRepository) FindCount() (int, error) {
	count := 0
	if err := d.db.Table(models.DingdingTableName).Where("IsDeleted=0").Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "find all dingding error")
	}
	return count, nil
}

func (d dingdingRepository) Update(dingding *models.Dingding) error {
	dingding.DataChange_CreatedTime = time.Now()
	db := d.db.Begin()
	if err := db.Table(models.DingdingTableName).Update(&dingding).Error; err != nil {
		db.Rollback()
		return errors.Wrap(err, "update dingding error")
	}
	db.Commit()
	return nil
}

func (d dingdingRepository) Delete(id int) error {
	db := d.db.Begin()
	if err := db.Table(models.DingdingTableName).Where("Id=?", id).Update("IsDeleted", 1).Error; err != nil {
		db.Rollback()
		return errors.Wrap(err, "delete dingding error")
	}
	db.Commit()
	return nil
}

func (d dingdingRepository) Find(Type, deptName, env string, level int) (*models.Dingding, error) {
	dingding := new(models.Dingding)
	if err := d.db.Table(models.DingdingTableName).First(&dingding, "Type=? and DeptName=? and Level=? and Env=? and IsDeleted=0", Type, deptName, level, env).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "find dingding error")
	}
	return dingding, nil
}

//查询名字是否重复
func (d dingdingRepository) FindByName(name string) bool {
	dingding := new(models.Dingding)
	if err := d.db.Table(models.DingdingTableName).First(&dingding, "Name=? and IsDeleted=0", name).Error; err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}
