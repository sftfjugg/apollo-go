package repositories

import (
	"apollo-adminserivce/internal/app/portal/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

//通过创始人查询暂定
type AppdRepository interface {
	Create(db *gorm.DB, app *models.App) error
	Update(db *gorm.DB, app *models.App) error
	FindByName(name string) (*models.App, error)
	DeleteByAppId(db *gorm.DB, appId string) error
	FindByAppId(appId string) (*models.App, error)
}

type appRepository struct {
	db *gorm.DB
}

func NewAppRepository(db *gorm.DB) AppdRepository {
	return &appRepository{db: db}
}

func (r appRepository) Create(db *gorm.DB, app *models.App) error {
	app.DataChange_LastTime = time.Now()
	app.DataChange_CreatedTime = time.Now()
	if err := db.Create(app).Error; err != nil {
		return errors.Wrap(err, "create app error")
	}
	return nil
}

func (r appRepository) Update(db *gorm.DB, app *models.App) error {
	app.DataChange_LastTime = time.Now()
	if err := db.Save(app).Error; err != nil {
		return errors.Wrap(err, "update app error")
	}
	return nil
}

func (r appRepository) FindByName(name string) (*models.App, error) {
	var app = new(models.App)
	if err := r.db.Table(models.AppTableName).First(&app, "(Name=?) and IsDeleted=0", name).Limit(1).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "find app by name  error")
	}
	return app, nil
}

func (r appRepository) FindByAppId(appId string) (*models.App, error) {
	var app = new(models.App)
	if err := r.db.Table(models.AppTableName).First(&app, "(AppId=?) and IsDeleted=0", appId).Limit(1).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "find app by AppId error")
	}
	return app, nil
}

//逻辑删除，并非实际删除
func (r appRepository) DeleteByAppId(db *gorm.DB, appId string) error {
	if err := db.Table(models.AppTableName).Where("AppId", appId).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "delete app by AppId error")
	}
	return nil
}
