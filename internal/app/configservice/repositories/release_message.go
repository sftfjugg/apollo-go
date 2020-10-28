package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type ReleaseMessageRepository interface {
	FindAll() ([]*models.ReleaseMessage, error)
	FindOffsetByMax(max int) ([]*models.ReleaseMessage, error)
}

type releaseMessageRepository struct {
	db *gorm.DB
}

func NewReleaseMessageRepository(db *gorm.DB) ReleaseMessageRepository {
	return &releaseMessageRepository{db: db}
}

func (r releaseMessageRepository) FindAll() ([]*models.ReleaseMessage, error) {
	var releases = make([]*models.ReleaseMessage, 0)
	if err := r.db.Table(models.ReleaseMessageTableName).Find(&releases).Error; err != nil {
		return nil, errors.Wrap(err, "find releaseMessage all error")
	}
	return releases, nil
}

func (r releaseMessageRepository) FindOffsetByMax(max int) ([]*models.ReleaseMessage, error) {
	var releases = make([]*models.ReleaseMessage, 0)
	if err := r.db.Table(models.ReleaseMessageTableName).Find(&releases, "id>?", max).Error; err != nil {
		return nil, errors.Wrap(err, "find releaseMessage all error")
	}
	return releases, nil
}
