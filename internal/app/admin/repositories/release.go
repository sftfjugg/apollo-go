package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type Release interface {
	Create(db *gorm.DB, release *models.Release) error
	Delete(db *gorm.DB, appId, cluster, namespace string) error
}

type release struct {
	db *gorm.DB
}

func NewRelease(db *gorm.DB) Release {
	return &release{
		db: db,
	}
}

func (r release) Create(db *gorm.DB, release *models.Release) error {
	release.DataChange_CreatedTime = time.Now()
	release.DataChange_LastTime = time.Now()
	if err := db.Create(release).Error; err != nil {
		return errors.Wrap(err, "create release error")
	}
	return nil
}

//物理删除，保持发布数目一直稳定，以防止数据量累积导致速度过慢
func (r release) Delete(db *gorm.DB, appId, cluster, namespace string) error {
	if err := db.Table(models.ReleaseTableName).Delete(&models.Release{}, "AppId= ? and Cluster=? and Namespace=? ", appId, cluster, namespace).Error; err != nil {
		return errors.Wrap(err, "delete previous release error")
	}
	return nil
}
