package repositories

import (
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type AppNamespaceRepository interface {
	Create(db *gorm.DB, AppNamespace *models.AppNamespace) error
	Delete(db *gorm.DB, id string) error
	Update(db *gorm.DB, AppNamespace *models.AppNamespace) error
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
	FindAppNamespaceIsPublic(appId string) ([]*models.AppNamespace, error)
}

type appNamespaceRepository struct {
	db *gorm.DB
}

func NewAppNamespaceRepository(db *gorm.DB) AppNamespaceRepository {
	return &appNamespaceRepository{
		db: db,
	}
}

func (r appNamespaceRepository) Create(db *gorm.DB, appNamespace *models.AppNamespace) error {
	if err := db.Create(appNamespace).Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appNamespaceRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Table(models.AppNamespaceTableName).Update("IsDeleted=?", true).Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appNamespaceRepository) Update(db *gorm.DB, appNamespace *models.AppNamespace) error {
	if err := db.Save(appNamespace).Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appNamespaceRepository) FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(appNamespaces, "AppId=? and clusterName=? and and IsPublic=1", appId, clusterName).Error; err != nil {
		return nil, errors.Wrap(err, "create app error")
	}
	return appNamespaces, nil
}

func (r appNamespaceRepository) FindAppNamespaceIsPublic(appId string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(appNamespaces, "AppId=? and IsPublic=1?", appId).Error; err != nil {
		return nil, errors.Wrap(err, "create app error")
	}
	return appNamespaces, nil
}
