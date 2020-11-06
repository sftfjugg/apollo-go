package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type ReleaseHistoryRepository interface {
	Create(db *gorm.DB, releaseHistory *models.ReleaseHistory) error
	Find(appId, namespaceName, key string, pageSize, pageNum int) ([]*models.ReleaseHistory, error)
	FindCount(appId, namespaceName, key string) (int, error)
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
		key = "and OperationContext like '%" + key + "%'"
	}
	releaseHistorys := make([]*models.ReleaseHistory, 0)
	if err := r.db.Table(models.ReleaseHistoryTableName).Limit(pageSize).Offset(pageSize*(pageNum-1)).Find(&releaseHistorys, "AppId=? and NamespaceName=? "+key, appId, namespaceName).Error; err != nil {
		return nil, errors.Wrap(err, "find releaseHistory error")
	}
	return releaseHistorys, nil
}

func (r releaseHistoryRepository) FindCount(appId, namespaceName, key string) (int, error) {
	if key != "" {
		key = "and Key OperationContext '%" + key + "%'"
	}
	var count = new(models2.Count)
	if err := r.db.Raw("Select count(*) as count  from `ReleaseHistory`  where AppId=? and NamespaceName=?"+key, appId, namespaceName).Scan(&count).Error; err != nil {
		return 0, errors.Wrap(err, "ItemRepisitory.FindSum failed")
	}
	return count.Count, nil
}
