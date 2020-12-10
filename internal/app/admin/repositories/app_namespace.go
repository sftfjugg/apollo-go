package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type AppNamespaceRepository interface {
	Create(db *gorm.DB, appNamespace *models.AppNamespace) error
	DeleteById(db *gorm.DB, id string) error
	DeleteByNameAndAppIdAndCluster(db *gorm.DB, name, appId, cluster string) error
	//Update(db *gorm.DB, appNamespace *models.AppNamespace) error
	Update(db *gorm.DB, appNamespace *models.AppNamespace) error
	FindAppNamespaceById(id string) (*models.AppNamespace, error)
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
	FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appId, clusterName, laneName, name string) (*models.AppNamespace, error)
	FindAppNamespace(appId, cluster, format string) ([]*models.AppNamespace, error)
	FindAppNamespaceByAppIdAndName(appId, name string) ([]*models.AppNamespace, error)
	FindClusterNameByAppId(appId string) ([]*models.AppNamespace, error)
	FindAllClusterNameByAppId(appId string) ([]*models.AppNamespace, error)
	FindByLaneName(lane string) ([]*models.AppNamespace, error)
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
func (r appNamespaceRepository) DeleteByNameAndAppIdAndCluster(db *gorm.DB, name, appId, cluster string) error {
	if err := db.Table(models.AppNamespaceTableName).Where("Name=? and AppId=? and ClusterName=?", name, appId, cluster).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "delete appNamespace by appid and name error")
	}
	return nil
}

//func (r appNamespaceRepository) Update(db *gorm.DB, appNamespace *models.AppNamespace) error {
//	appNamespace.DataChange_LastTime = time.Now()
//	if err := db.Table(models.AppNamespaceTableName).Where("Id=?", appNamespace.Id).Update(appNamespace).Error; err != nil {
//		return errors.Wrap(err, "update appNamespace error")
//	}
//	return nil
//}

func (r appNamespaceRepository) Update(db *gorm.DB, appNamespace *models.AppNamespace) error {
	appNamespace.DataChange_LastTime = time.Now()
	if err := db.Exec(" UPDATE `AppNamespace` SET  `Comment` = ?, `DataChange_LastModifiedBy` = ?, `DataChange_LastTime` = ?, `DeptName` = ?, `Format` = ?, `IsDisplay` = ?  , `IsPublic` = ?  WHERE (Name=? and AppId=? and ClusterName=? and IsDeleted=0)", appNamespace.Comment, appNamespace.DataChange_LastModifiedBy, time.Now(), appNamespace.DeptName, appNamespace.Format, appNamespace.IsDisplay, appNamespace.IsPublic, appNamespace.Name, appNamespace.AppId, appNamespace.ClusterName).Error; err != nil {
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

func (r appNamespaceRepository) FindAppNamespaceById(id string) (*models.AppNamespace, error) {
	appNamespace := new(models.AppNamespace)
	if err := r.db.Table(models.AppNamespaceTableName).First(&appNamespace, "Id=?  and IsDeleted=0", id).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceById appNamespace error")
	}
	return appNamespace, nil
}

func (r appNamespaceRepository) FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appId, clusterName, laneName, name string) (*models.AppNamespace, error) {
	appNamespace := new(models.AppNamespace)
	if err := r.db.Table(models.AppNamespaceTableName).First(&appNamespace, "AppId=? and ClusterName=? and name=? and laneName=?  and IsDeleted=0", appId, clusterName, name, laneName).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane appNamespace error")
	}
	return appNamespace, nil
}

//查询appId下的所有配置
func (r appNamespaceRepository) FindAppNamespace(appId, cluster, format string) ([]*models.AppNamespace, error) {
	if format != "" {
		format = "and Format='" + format + "'  "
	}
	if cluster != "" {
		cluster = "and ClusterName='" + cluster + "' "
	}
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "AppId=? and IsDeleted=0 "+format+cluster+" ", appId).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespace appNamespace error")
	}
	return appNamespaces, nil
}

//查询appId下的所有集群名字
func (r appNamespaceRepository) FindClusterNameByAppId(appId string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Raw("select * from AppNamespace where AppId=? and IsDeleted=0;", appId).Scan(&appNamespaces).Error; err != nil {
		return nil, errors.Wrap(err, "FindClusterNameByAppId appNamespace error")
	}
	return appNamespaces, nil
}

func (r appNamespaceRepository) FindAppNamespaceByAppIdAndName(appId, name string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if err := r.db.Table(models.AppNamespaceTableName).Find(&appNamespaces, "AppId=? and IsDeleted=0 and Name=?", appId, name).Error; err != nil {
		return nil, errors.Wrap(err, "FindAppNamespaceByAppIdAndName appNamespace error")
	}
	return appNamespaces, nil
}

func (r appNamespaceRepository) FindAllClusterNameByAppId(appId string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if appId != "" {
		appId = "and AppId=" + "'" + appId + "' "
	}
	if err := r.db.Raw(" select ClusterName FROM `AppNamespace`  WHERE IsDeleted=0 " + appId + "  group by ClusterName").Scan(&appNamespaces).Error; err != nil {
		return nil, errors.Wrap(err, "FindAllClusterNameByAppId appNamespace error")
	}
	return appNamespaces, nil
}

func (r appNamespaceRepository) FindByLaneName(lane string) ([]*models.AppNamespace, error) {
	appNamespaces := make([]*models.AppNamespace, 0)
	if lane != "" {
		lane = "and LaneName='" + lane + "'"
	}
	if err := r.db.Raw(" select * FROM `AppNamespace`  WHERE IsDeleted=0 " + lane).Scan(&appNamespaces).Error; err != nil {
		return nil, errors.Wrap(err, "FindByLaneName appNamespace error")
	}
	return appNamespaces, nil
}
