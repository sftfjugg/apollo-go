package services

import (
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
	FindGroupsOfDevelopment() (*uic_service_api.Node, error)
	FindLimosAppForPage(name string, pageSize, pageNum int32) (*plat_limos_rpc.AppsForPage, error)
	GetAllUsers(name string) ([]*uic_service_api.UserData, error)
	FindLimosAppById(appID int64) (*plat_limos_rpc.App, error)
	FindAuth(appId int64, name string) (bool, error)
}

type appService struct {
	limosService plat_limos_rpc.TChanLimosService
	uicService   uic_service_api.TChanUicService
}

func NewAppService(
	limosService plat_limos_rpc.TChanLimosService,
	uicService uic_service_api.TChanUicService,
) AppService {
	return &appService{
		limosService: limosService,
		uicService:   uicService,
	}
}

func (s appService) FindGroupsOfDevelopment() (*uic_service_api.Node, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	groups, err := s.uicService.FindGroupsOfDevelopment(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.FindGroupsOfDevelopment() error")
	}
	return groups, nil
}

func (s appService) GetAllUsers(name string) ([]*uic_service_api.UserData, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	names, err := s.uicService.GetAllUsers(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.GetAllUsers() error")
	}
	return names, nil
}

func (s appService) FindLimosAppForPage(name string, pageSize, pageNum int32) (*plat_limos_rpc.AppsForPage, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	apps, err := s.limosService.FindAppForPage(ctx, name, "", "", "", "", 0, "all", pageSize, pageNum, "", 0)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.FindAppForPage() error")
	}
	return apps, nil
}

func (s appService) FindLimosAppById(appID int64) (*plat_limos_rpc.App, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	app, err := s.limosService.GetAppByID(ctx, appID)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.GetAppByID() error")
	}
	return app, nil
}

func (s appService) FindAuth(appId int64, name string) (bool, error) {
	app, err := s.FindLimosAppById(appId)
	if err != nil {
		return false, errors.Wrap(err, "call zclients uic.GetAppByID() error")
	}
	for o := range app.Owner {
		if strings.Compare(app.Owner[o], name) == 0 {
			return true, nil
		}
	}
	return false, nil
}
