package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/models"
)

type ConfigRepository interface {
	FindPublicConfig(appId string) ([]*models.Config, error)
	FindPrivateConfig(appId, cluster string) ([]*models.Config, error)
	FindConfig(appId, cluster, namespace string) ([]*models.Config, error)
}

type configRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) ConfigRepository {
	return &configRepository{db: db}
}

func (r configRepository) FindPublicConfig(appId string) ([]*models.Config, error) {
	var configurations = make([]*models.Config, 0)
	if err := r.db.Raw("select Configurations from `Release`  where ClusterName='default' and IsDeleted=0 and Id in(select max(Id) from `Release`  group by AppId,NamespaceName having AppId=?) ", appId).Scan(&configurations).Error; err != nil {
		return nil, errors.Wrap(err, "find config publish  error")
	}
	return configurations, nil
}

func (r configRepository) FindPrivateConfig(appId, cluster string) ([]*models.Config, error) {
	var configurations = make([]*models.Config, 0)
	if err := r.db.Raw("select  `AppId`,`ReleaseKey`,`ClusterName`,`NamespaceName`,`Configurations` from `Release` where ClusterName=? and IsDeleted=0  and Id in (select max(Id) from `Release`  group by AppId,NamespaceName having AppId=?)", cluster, appId).Scan(&configurations).Error; err != nil {
		return nil, errors.Wrap(err, "find config private  error")
	}
	return configurations, nil
}

func (r configRepository) FindConfig(appId, cluster, namespcae string) ([]*models.Config, error) {
	var configurations = make([]*models.Config, 0)
	if err := r.db.Raw("select  `AppId`,`ReleaseKey`,`ClusterName`,`NamespaceName`,`Configurations` from `Release` where AppId=? and ClusterName=? and IsDeleted=0  and  NamespaceName=? order by Id desc limit 1", appId, cluster, namespcae).Scan(&configurations).Error; err != nil {
		return nil, errors.Wrap(err, "find config private  error")
	}
	return configurations, nil
}
