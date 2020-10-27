package repositories

import (
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

type AppNamespaceRepository interface {
	Create(db *gorm.DB, appNamespace *models.AppNamespace) error
	DeleteById(db *gorm.DB, id string) error
	DeleteByNameAndAppId(db *gorm.DB, name, appId string) error
	Update(db *gorm.DB, appNamespace *models.AppNamespace) error
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
	FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name string) (*models.AppNamespace, error)
	FindAppNamespaceByAppId(appId, format string) ([]*models.AppNamespace, error)
	FindAppNamespaceByAppIdAndName(appId, name string) ([]*models.AppNamespace, error)
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
	appNamespace.DataChange_LastTime = time.Now()
	appNamespace.DataChange_CreatedTime = time.Now()
	if err := db.Create(&appNamespace).Error; err != nil {
		return errors.Wrap(err, "create appNamespace error")
	}
	return nil
}

func (r appNamespaceRepository) DeleteById(db *gorm.DB, id string) error {
	if err := db.Table(models.AppNamespaceTableName).Where("Id=?", id).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "delete appNamespace by id error")
	}
	return nil
}

//多表删除，通过配置文件名字删除掉配置文件和对应配置
func (r appNamespaceRepository) DeleteByNameAndAppId(db *gorm.DB, name, appId string) error {
	if err := db.Table(models.AppNamespaceTableName).Where("Name=? and AppId=?", name, appId).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "delete appNamespace by appid and name error")
	}
	return nil
}

func (r appNamespaceRepository) Update(db *gorm.DB, appNamespace *models.AppNamespace) error {
	appNamespace.DataChange_LastTime = time.Now()
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

func (r appNamespaceRepository) FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name string) (*models.AppNamespace, error) {
	appNamespace := new(models.AppNamespace)
	if err := r.db.Table(models.AppNamespaceTableName).First(&appNamespace, "AppId=? and ClusterName=? and name=?  and IsDeleted=0", appId, clusterName, name).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "FindOneAppNamespaceByAppIdAndClusterName appNamespace error")
	}
	return appNamespace, nil
}

//查询appId下的所有配置
func (r appNamespaceRepository) FindAppNamespaceByAppId(appId, format string) ([]*models.AppNamespace, error) {
	if format != "" {
		format = "and Format='" + format + "'"
	}
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "AppId=? and IsDeleted=0 "+format+"", appId).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceByAppId appNamespace error")
	}
	return appNamespaces, nil
}

//查询appId下的所有集群名字
func (r appNamespaceRepository) FindClusterNameByAppId(appId string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Raw("select ClusterName from AppNamespace where AppId=? and IsDeleted=0 group by ClusterName order by null;", appId).Scan(&appNamespaces).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceIsPublic appNamespace error")
	}
	return appNamespaces, nil
}

func (r appNamespaceRepository) FindAppNamespaceByAppIdAndName(appId, name string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "AppId=? and IsDeleted=0 and Name=?", appId, name).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceByAppId appNamespace error")
	}
	return appNamespaces, nil
}
