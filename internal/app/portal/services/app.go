package services

import (
	"apollo-adminserivce/internal/app/portal/repositories"
	"apollo-adminserivce/internal/pkg/models"
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
	FindAllForPage(pageSize, pageNum int) ([]*models.App, error)
	FindByNameOrAppIdForPage(name, appId string, pageSize, pageNum int) ([]*models.App, error)
	FindByNameForPage(name string, pageSize, pageNum int) ([]*models.App, error)
	FindByAppId(appId string) (*models.App, error)
	DeleteByAppId(appId string) error
	FindGroupsOfDevelopment() (*uic_service_api.Node, error)
	FindLimosAppForPage(name string, pageSize, pageNum int32) (*plat_limos_rpc.AppsForPage, error)
	GetAllUsers(name string) ([]*uic_service_api.UserData, error)
	FindLimosAppById(appID int64) (*plat_limos_rpc.App, error)
	FindAuth(appId int64, name string) (bool, error)
}

type appService struct {
	repository   repositories.AppRepository
	db           *gorm.DB
	limosService plat_limos_rpc.TChanLimosService
	uicService   uic_service_api.TChanUicService
}

func NewAppService(
	repository repositories.AppRepository,
	db *gorm.DB,
	limosService plat_limos_rpc.TChanLimosService,
	uicService uic_service_api.TChanUicService,
) AppService {
	return &appService{
		repository:   repository,
		db:           db,
		limosService: limosService,
		uicService:   uicService,
	}
}

func (s appService) Create(app *models.App) error {
	db := s.db.Begin()
	if err := s.repository.Create(db, app); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s appService) Update(app *models.App) error {
	db := s.db.Begin()
	if err := s.repository.Update(db, app); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppRepository.Update() error")
	}
	db.Commit()
	return nil
}

func (s appService) FindAllForPage(pageSize, pageNum int) ([]*models.App, error) {
	apps, err := s.repository.FindAllForPage(pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call AppRepository.Update() error")
	}
	return apps, err
}

func (s appService) FindByNameOrAppIdForPage(name, appId string, pageSize, pageNum int) ([]*models.App, error) {
	apps, err := s.repository.FindByNameOrAppIdForPage(name, appId, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call AppRepository.FindByNameOrAppIdForPage() error")
	}
	return apps, err
}

func (s appService) FindByNameForPage(name string, pageSize, pageNum int) ([]*models.App, error) {
	apps, err := s.repository.FindByNameForPage(name, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call AppRepository.FindByNameForPage() error")
	}
	return apps, err
}

func (s appService) FindByAppId(appId string) (*models.App, error) {
	app, err := s.repository.FindByAppId(appId)
	if err != nil {
		return nil, errors.Wrap(err, "call AppRepository.FindByAppId() error")
	}
	return app, nil
}

func (s appService) DeleteByAppId(appId string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteByAppId(db, appId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call DeleteByAppId.Update() error")
	}
	db.Commit()
	return nil
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
