package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type FavoriteRepository interface {
	FindByUserIdForPage(UserId string, pageSize, pageNum int) ([]*models.App, error)
	Create(db *gorm.DB, favorite *models.Favorite) error
	DeleteByUserIdAndAppId(db *gorm.DB, userId, appId string) error
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteReposity(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{
		db: db,
	}
}

func (r favoriteRepository) FindByUserIdForPage(UserId string, pageSize, pageNum int) ([]*models.App, error) {
	var apps = make([]*models.App, 0)
	if err := r.db.Raw("select App.AppId,App.Name from App,Favorite where App.AppId=Favorite.AppId and Favorite.UserId=? and App.IsDeleted=Favorite.IsDeleted and App.IdDeleted=0", UserId).Limit(pageSize).Offset(pageSize * (pageNum - 1)).Scan(&apps).Error; err != nil {
		return nil, errors.Wrap(err, "find Favorite all error")
	}
	return apps, nil
}

func (r favoriteRepository) Create(db *gorm.DB, favorite *models.Favorite) error {
	if err := db.Create(favorite).Error; err != nil {
		return errors.Wrap(err, "create favorite error")
	}
	return nil
}

func (r favoriteRepository) DeleteByUserIdAndAppId(db *gorm.DB, userId, appId string) error {
	if err := db.Table(models.FavoriteTableName).Where("UserId=? and AppId=?", userId, appId).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "delete favorite by AppId error")
	}
	return nil
}
