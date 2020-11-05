package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type ReleaseHistoryRepository interface {
	Create(db *gorm.DB, releaseHistory *models.ReleaseHistory) error
	Find(appId, namespaceName, key string, pageSize, pageNum int) ([]*models.ReleaseHistory, error)
}

type releaseHistoryRepository struct {
	db *gorm.DB
}

func NewReleaseHistoryRepository(db *gorm.DB) ReleaseHistoryRepository {
	return &releaseHistoryRepository{db: db}
}

func (r releaseHistoryRepository) Create(db *gorm.DB, releaseHistory *models.ReleaseHistory) error {
	releaseHistory.DataChange_CreatedTime = time.Now()
	releaseHistory.DataChange_LastTime = time.Now()
	if err := db.Create(releaseHistory).Error; err != nil {
		return errors.Wrap(err, "create releaseHistory error")
	}
	return nil
}

func (r releaseHistoryRepository) Find(appId, namespaceName, key string, pageSize, pageNum int) ([]*models.ReleaseHistory, error) {
	if key != "" {
		key = "and Key like '%" + key + "%'"
	}
	releaseHistorys := make([]*models.ReleaseHistory, 0)
	if err := r.db.Table(models.ReleaseHistoryTableName).Limit(pageSize).Offset(pageSize*(pageNum-1)).Find(&releaseHistorys, "AppId=? and NamespaceName=? "+key, appId, namespaceName).Error; err != nil {
		return nil, errors.Wrap(err, "find releaseHistory error")
	}
	return releaseHistorys, nil
}
