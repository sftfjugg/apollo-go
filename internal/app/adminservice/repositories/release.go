package repositories

import (
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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
	if err := db.Create(release).Error; err != nil {
		return errors.Wrap(err, "create release error")
	}
	return nil
}
