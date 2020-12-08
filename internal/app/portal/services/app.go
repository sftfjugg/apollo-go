package services

import (
	"github.com/pkg/errors"
	"github.com/uber/tchannel-go"
	"go.didapinche.com/goapi/plat_limos_rpc"
	"go.didapinche.com/goapi/uic_service_api"
	"time"
)

//apps,_:=app.FindGroupsOfDevelopment(ctx)获得全部组
//apps,_:=app.FindGroupsOfDevelopment(ctx)获得所有用户

type AppService interface {
	FindGroupsOfDevelopment() (*uic_service_api.Node, error)
	FindLimosAppForPage(name, owner string, pageSize, pageNum int32) (*plat_limos_rpc.AppsForPage, error)
	GetAllUsers(name string) ([]*uic_service_api.UserData, error)
	FindAuth(appId, userId string) (int, error)
}

type appService struct {
	limosService plat_limos_rpc.TChanLimosService
	uicService   uic_service_api.TChanUicService
	roleService  RoleService
}

func NewAppService(
	limosService plat_limos_rpc.TChanLimosService,
	uicService uic_service_api.TChanUicService,
	roleService RoleService,
) AppService {
	return &appService{
		limosService: limosService,
		uicService:   uicService,
		roleService:  roleService,
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

func (s appService) FindLimosAppForPage(name, owner string, pageSize, pageNum int32) (*plat_limos_rpc.AppsForPage, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	apps, err := s.limosService.FindAppForPage(ctx, name, "", "", owner, "", 0, "all", pageSize, pageNum, "", 0)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.FindAppForPage() error")
	}
	return apps, nil
}

func (s appService) FindAuth(appId, userId string) (int, error) {

	//验证owner权限
	app, err := s.FindLimosAppForPage(appId, userId, 20, 1)
	if err != nil {
		return 0, errors.Wrap(err, "call zclients uic.GetAppByID() error")
	}
	if app.TotalCount > 0 {
		return 4, nil
	}

	i, err := s.roleService.Find(appId, userId)
	if err != nil {
		return 0, errors.Wrap(err, "call RoleService.Find error")
	}
	return i, nil
}
