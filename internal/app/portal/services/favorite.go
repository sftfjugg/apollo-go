package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type FavoriteService interface {
	FindByUserIdForPage(UserId string, pageSize, pageNum int) ([]*models.App, error)
	Create(favorite *models.Favorite) error
	DeleteByUserIdAndAppId(userId, appId string) error
}

type favoriteService struct {
	repository repositories.FavoriteRepository
	db         *gorm.DB
}

func NewFavoriteService(repository repositories.FavoriteRepository, db *gorm.DB) FavoriteService {
	return &favoriteService{
		db:         db,
		repository: repository,
	}
}

func (s favoriteService) FindByUserIdForPage(UserId string, pageSize, pageNum int) ([]*models.App, error) {
	app, err := s.repository.FindByUserIdForPage(UserId, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call FavoriteRepository.FindByUserIdForPage() error")
	}
	return app, nil
}

func (f favoriteService) Create(favorite *models.Favorite) error {
	db := f.db.Begin()
	if err := f.repository.Create(db, favorite); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call FavoriteRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (f favoriteService) DeleteByUserIdAndAppId(userId, appId string) error {
	db := f.db.Begin()
	if err := f.repository.DeleteByUserIdAndAppId(db, userId, appId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call FavoriteRepository.DeleteByUserId() error")
	}
	db.Commit()
	return nil
}
