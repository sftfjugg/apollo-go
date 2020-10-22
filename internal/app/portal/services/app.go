package services

import (
	"apollo-adminserivce/internal/app/portal/models"
	"apollo-adminserivce/internal/app/portal/repositories"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/uber/tchannel-go"
	"go.didapinche.com/goapi/plat_limos_rpc"
	"go.didapinche.com/goapi/uic_service_api"
	"strings"
	"time"
)

//apps,_:=app.FindGroupsOfDevelopment(ctx)获得全部组
//apps,_:=app.FindGroupsOfDevelopment(ctx)获得所有用户

type AppService interface {
	Create(app *models.App) error
	Update(app *models.App) error
	FindByName(name string) (*models.App, error)
	FindByAppId(appId string) (*models.App, error)
	DeleteByAppId(appId string) error
	FindGroupsOfDevelopment() (*uic_service_api.Node, error)
	FindLimosAppForPage(name string, pageSize, pageNum int32) (*plat_limos_rpc.AppsForPage, error)
	GetAllUsers(name string) ([]*uic_service_api.UserData, error)
	FindLimosAppById(appID int64) (*plat_limos_rpc.App, error)
	FindAuth(appId int64, name string) (bool, error)
}

type appService struct {
	db           *gorm.DB
	repository   repositories.AppdRepository
	limosService plat_limos_rpc.TChanLimosService
	uicService   uic_service_api.TChanUicService
}

func NewAppService(
	db *gorm.DB,
	limosService plat_limos_rpc.TChanLimosService,
	uicService uic_service_api.TChanUicService,
	repository repositories.AppdRepository,
) AppService {
	return &appService{
		db:           db,
		repository:   repository,
		limosService: limosService,
		uicService:   uicService,
	}
}

func (s appService) Create(app *models.App) error {
	apps, err := s.FindByAppId(app.AppId)
	if err != nil {
		return errors.Wrap(err, "call AppRepository.FindByName() error")
	}
	if apps.AppId != "" {
		return errors.New("Appid already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Create(s.db, app); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s appService) Update(app *models.App) error {
	apps, err := s.FindByAppId(app.AppId)
	if err != nil {
		return errors.Wrap(err, "call AppRepository.FindByName() error")
	}
	if apps.AppId != "" || apps.AppId != app.AppId {
		return errors.New("Appid already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Update(s.db, app); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppRepository.Update() error")
	}
	db.Commit()
	return nil
}

func (s appService) DeleteByAppId(appId string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteByAppId(s.db, appId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppRepository.DeleteByAppId() error")
	}
	db.Commit()
	return nil
}

func (s appService) FindByName(name string) (*models.App, error) {
	app, err := s.repository.FindByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "call AppRepository.FindByName() error")
	}
	return app, nil
}

func (s appService) FindByAppId(appId string) (*models.App, error) {
	app, err := s.repository.FindByAppId(appId)
	if err != nil {
		return nil, errors.Wrap(err, "call AppRepository.FindByAppId() error")
	}
	return app, nil
}

func (s appService) FindGroupsOfDevelopment() (*uic_service_api.Node, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	groups, err := s.uicService.FindGroupsOfDevelopment(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "call clients uic.FindGroupsOfDevelopment() error")
	}
	return groups, nil
}

func (s appService) GetAllUsers(name string) ([]*uic_service_api.UserData, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	names, err := s.uicService.GetAllUsers(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "call clients uic.GetAllUsers() error")
	}
	return names, nil
}

func (s appService) FindLimosAppForPage(name string, pageSize, pageNum int32) (*plat_limos_rpc.AppsForPage, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	apps, err := s.limosService.FindAppForPage(ctx, name, "", "", "", "", 0, "all", pageSize, pageNum, "", 0)
	if err != nil {
		return nil, errors.Wrap(err, "call clients uic.FindAppForPage() error")
	}
	return apps, nil
}

func (s appService) FindLimosAppById(appID int64) (*plat_limos_rpc.App, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	app, err := s.limosService.GetAppByID(ctx, appID)
	if err != nil {
		return nil, errors.Wrap(err, "call clients uic.GetAppByID() error")
	}
	return app, nil
}

func (s appService) FindAuth(appId int64, name string) (bool, error) {
	app, err := s.FindLimosAppById(appId)
	if err != nil {
		return false, errors.Wrap(err, "call clients uic.GetAppByID() error")
	}
	for o := range app.Owner {
		if strings.Compare(app.Owner[o], name) == 0 {
			return true, nil
		}
	}
	return false, nil
}
