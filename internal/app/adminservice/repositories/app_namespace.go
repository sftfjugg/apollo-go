package repositories

import (
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type AppNamespaceRepository interface {
	Create(db *gorm.DB, appNamespace *models.AppNamespace) error
	DeleteById(db *gorm.DB, id string) error
	Update(db *gorm.DB, appNamespace *models.AppNamespace) error
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
	FindAppNamespaceByIsPublic(appId string) ([]*models.AppNamespace, error)
	FindClusterNameByAppId(appId string) ([]*models.AppNamespace, error)
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
		return errors.Wrap(err, "create appNamespace error")
	}
	return nil
}

func (r appNamespaceRepository) DeleteById(db *gorm.DB, id string) error {
	if err := db.Table(models.AppNamespaceTableName).Where("Id=?", id).Update("IsDeleted=?", true).Error; err != nil {
		return errors.Wrap(err, "delete appNamespace error")
	}
	return nil
}

func (r appNamespaceRepository) Update(db *gorm.DB, appNamespace *models.AppNamespace) error {
	if err := db.Table(models.AppNamespaceTableName).Where("Id=?", appNamespace.Id).Update(appNamespace).Error; err != nil {
		return errors.Wrap(err, "update appNamespace error")
	}
	return nil
}

func (r appNamespaceRepository) FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "AppId=? and ClusterName=?  and IsDeleted=0", appId, clusterName).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceByAppIdAndClusterName appNamespace error")
	}
	return appNamespaces, nil
}

//查询appId下的公共配置的名字和集群名字
func (r appNamespaceRepository) FindAppNamespaceByIsPublic(appId string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "AppId=? and IsPublic=1? and IsDeleted=0", appId).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceIsPublic appNamespace error")
	}
	return appNamespaces, nil
}

//查询appId下的所有集群名字
func (r appNamespaceRepository) FindClusterNameByAppId(appId string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Select("ClusterName").Find(&appNamespaces, "AppId=? and IsDeleted=0", appId).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceIsPublic appNamespace error")
	}
	return appNamespaces, nil
}
