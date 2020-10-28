package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice/models"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type ConfigRepository interface {
	FindPublicConfigName(appId string) ([]*models2.AppNamespace, error)
	FindPublicConfig(appId, name string) ([]*models.Config, error)
	FindPrivateConfig(appId, cluster string) ([]*models.Config, error)
}

type configRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) ConfigRepository {
	return &configRepository{db: db}
}

func (r configRepository) FindPublicConfigName(appId string) ([]*models2.AppNamespace, error) {
	var names = make([]*models2.AppNamespace, 0)
	if err := r.db.Table(models2.AppNamespaceTableName).Select("Name").Find(&names, "appId=? and IsDeleted=0 and IsPublic=1", appId).Error; err != nil {
		return nil, errors.Wrap(err, "find configName  error")
	}
	return names, nil
}

func (r configRepository) FindPublicConfig(appId, name string) ([]*models.Config, error) {
	var configurations = make([]*models.Config, 0)
	if err := r.db.Raw("select Configurations from `Release` where AppId=? and NamespaceName=? order by Id desc limit 1", appId, name).Scan(&configurations).Error; err != nil {
		return nil, errors.Wrap(err, "find config publish  error")
	}
	return configurations, nil
}

func (r configRepository) FindPrivateConfig(appId, cluster string) ([]*models.Config, error) {
	var configurations = make([]*models.Config, 0)
	if err := r.db.Raw("select  `AppId`,`ReleaseKey`,`ClusterName`,`NamespaceName`,`Configurations` from `Release` where AppId=? and ClusterName=? order by Id desc limit 1", appId, cluster).Scan(&configurations).Error; err != nil {
		return nil, errors.Wrap(err, "find config private  error")
	}
	return configurations, nil
}
