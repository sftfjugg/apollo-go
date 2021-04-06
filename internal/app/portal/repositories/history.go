package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type HistoryRepository interface {
	Create(history *models.History) error
	FindHistory(userId string) ([]*models.History, error)
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

func (h historyRepository) Create(history *models.History) error {
	history.DataChange_CreatedTime = time.Now()
	db := h.db.Begin()
	if err := db.Table(models.HistoryTableName).Create(&history).Error; err != nil {
		db.Rollback()
		return errors.Wrap(err, "create history error")
	}
	db.Commit()
	return nil
}

func (h historyRepository) FindHistory(userId string) ([]*models.History, error) {
	historys := make([]*models.History, 0)
	if err := h.db.Table(models.HistoryTableName).Limit(20).Order("Id desc").Group("AppId").Find(&historys, "UserId=?", userId).Error; err != nil {
		return nil, errors.Wrap(err, "call FindHistory error")
	}
	return historys, nil
}
