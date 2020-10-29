package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type Release interface {
	Create(db *gorm.DB, release *models.Release) error
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
