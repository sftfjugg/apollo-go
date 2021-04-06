package repositories

import (
	"github.com/jinzhu/gorm"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type DingdingRepository interface {
	Create(dingding *models.Dingding) error
	FindHistory(userId string) ([]*models.Dingding, error)
}

type dingdingRepository struct {
	db *gorm.DB
}
