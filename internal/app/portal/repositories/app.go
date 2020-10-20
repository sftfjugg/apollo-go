package repositories

import (
	"apollo-adminserivce/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//通过创始人查询暂定
type AppRepository interface {
	Create(db *gorm.DB, app *models.App) error
	Update(db *gorm.DB, app *models.App) error
	FindAllForPage(pageSize, pageNum int) ([]*models.App, error)
	FindByNameOrAppIdForPage(name, appId string, pageSize, pageNum int) ([]*models.App, error)
	FindByNameForPage(name string, pageSize, pageNum int) ([]*models.App, error)
	FindByAppId(appId string) (*models.App, error)
	DeleteByAppId(db *gorm.DB, appId string) error
}

type appRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) AppRepository {
	return &appRepository{db: db}
}

func (r appRepository) Create(db *gorm.DB, app *models.App) error {
	if err := db.Create(app).Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appRepository) Update(db *gorm.DB, app *models.App) error {
	if err := db.Save(app).Error; err != nil {
		return errors.Wrap(err, "update app error")
	}
	return nil
}

func (r appRepository) FindAllForPage(pageSize, pageNum int) ([]*models.App, error) {
	var apps = make([]*models.App, 0)
	if err := r.db.Table(models.AppTableName).Select("AppId", "Name").Find(apps, "IsDeleted=0").Limit(pageSize).Offset(pageSize * (pageNum - 1)).Error; err != nil {
		return nil, errors.Wrap(err, "find app all error")
	}
	return apps, nil
}

func (r appRepository) FindByNameOrAppIdForPage(name, appId string, pageSize, pageNum int) ([]*models.App, error) {
	var apps = make([]*models.App, 0)
	if err := r.db.Table(models.AppTableName).Select("AppId", "Name").Find(&apps, "(Name=? or AppId=?) and IsDeleted=0", name, appId).Limit(pageSize).Offset(pageSize * (pageNum - 1)).Error; err != nil {
		return nil, errors.Wrap(err, "find app by name of AppId error")
	}
	return apps, nil
}

func (r appRepository) FindByNameForPage(name string, pageSize, pageNum int) ([]*models.App, error) {
	var apps = make([]*models.App, 0)
	if err := r.db.Table(models.AppTableName).Select("AppId", "Name").Find(&apps, "(Name=?) and IsDeleted=0", name).Limit(pageSize).Offset(pageSize * (pageNum - 1)).Error; err != nil {
		return nil, errors.Wrap(err, "find app by name of AppId error")
	}
	return apps, nil
}

func (r appRepository) FindByAppId(appId string) (*models.App, error) {
	app := new(models.App)
	if err := r.db.Table(models.AppTableName).First(&app, "(AppId=?) and IsDeleted=0", appId).Error; err != nil {
		return nil, errors.Wrap(err, "find app by AppId of AppId error")
	}
	return app, nil
}

//逻辑删除，并非实际删除
func (r appRepository) DeleteByAppId(db *gorm.DB, appId string) error {
	if err := db.Table(models.AppTableName).Update("IsDeleted=?", true).Where("AppId", appId).Error; err != nil {
		return errors.Wrap(err, "delete app by AppId error")
	}
	return nil
}
