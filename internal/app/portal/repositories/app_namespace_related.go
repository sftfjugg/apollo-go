package repositories

import (
	"apollo-adminserivce/internal/app/portal/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

type AppNamespaceRelatedRepository interface {
	Create(db *gorm.DB, AppNamespace *models.AppNamespace) error
	Delete(db *gorm.DB, id string) error
	Update(db *gorm.DB, AppNamespace *models.AppNamespace) error
	FindAppNamespaceByName(name string) (*models.AppNamespace, error)
	FindAppNamespaceById(id string) (*models.AppNamespace, error)
	FindAppNamespaceByNameForPage(name string, pageSize, pageNum int) ([]*models.AppNamespace, error)
	FindAppNamespaceByDepartmentForPage(department string, pageSize, pageNum int) ([]*models.AppNamespace, error)
}

type appNamespaceRelatedRepository struct {
	db *gorm.DB
}

func NewAppNamespaceRelatedRepository(db *gorm.DB) AppNamespaceRelatedRepository {
	return &appNamespaceRelatedRepository{
		db: db,
	}
}

func (r appNamespaceRelatedRepository) Create(db *gorm.DB, appNamespace *models.AppNamespace) error {
	appNamespace.DataChange_LastTime = time.Now()
	appNamespace.DataChange_CreatedTime = time.Now()
	if err := db.Create(appNamespace).Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appNamespaceRelatedRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Table(models.AppNamespaceTableName).Where("Id =?", id).Update("IsDeleted=1").Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appNamespaceRelatedRepository) Update(db *gorm.DB, appNamespace *models.AppNamespace) error {
	appNamespace.DataChange_LastTime = time.Now()
	if err := db.Table(models.AppNamespaceTableName).Where("Id =?", appNamespace.Id).Update(appNamespace).Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appNamespaceRelatedRepository) FindAppNamespaceByNameForPage(name string, pageSize, pageNum int) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "Name like? and  IsDeleted=0", "%"+name+"%").Limit(pageSize).Offset(pageSize * (pageNum - 1)).Error; err != nil {
		return nil, errors.Wrap(err, "find app by name error")
	}
	return appNamespaces, nil
}

func (r appNamespaceRelatedRepository) FindAppNamespaceByName(name string) (*models.AppNamespace, error) {
	appNamespace := new(models.AppNamespace)
	if err := r.db.Table(models.AppNamespaceTableName).First(&appNamespace, "Name =? and  IsDeleted=0", name).Error; err != nil {
		return nil, errors.Wrap(err, "find app by name error")
	}
	return appNamespace, nil
}

func (r appNamespaceRelatedRepository) FindAppNamespaceById(id string) (*models.AppNamespace, error) {
	appNamespace := new(models.AppNamespace)
	if err := r.db.Table(models.AppNamespaceTableName).First(&appNamespace, "Id =? and  IsDeleted=0", id).Error; err != nil {
		return nil, errors.Wrap(err, "find app by id error")
	}
	return appNamespace, nil
}

func (r appNamespaceRelatedRepository) FindAppNamespaceByDepartmentForPage(department string, pageSize, pageNum int) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "Department=? and IsDeleted=0", department).Limit(pageSize).Offset(pageSize * (pageNum - 1)).Error; err != nil {
		return nil, errors.Wrap(err, "create app error")
	}
	return appNamespaces, nil
}
