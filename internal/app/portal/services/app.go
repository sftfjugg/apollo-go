package services

import (
	"github.com/pkg/errors"
	"github.com/uber/tchannel-go"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/constans"
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
	FindAuth(appId, userId, cluster, env string) (*models.Auth, error)
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

func (s appService) FindAuth(appId, userId, cluster, env string) (*models.Auth, error) {

	//验证root权限
	auth := new(models.Auth)
	b, err := s.AuthPerm(userId, constans.ApolloRoot)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.Auth() error")
	}
	if b {
		auth.IsRoot = true
	}

	//验证owner权限
	if appId != "public_global_config" {
		app, err := s.FindLimosAppForPage(appId, userId, 20, 0)
		if err != nil {
			return nil, errors.Wrap(err, "call zclients uic.GetAppByID() error")
		}
		if app.TotalCount > 0 {
			auth.IsOwner = true
		}
		//验证op权限
		o, err := s.AuthPerm(userId, constans.AppOperate)
		if err != nil {
			return nil, errors.Wrap(err, "call zclients uic.Auth() error")
		}
		if o {
			auth.IsOperate = true
		}
	} else {
		//公共配置权限默认为owner，不受op管理
		b, err := s.AuthPerm(userId, constans.ApolloPublicOperate)
		if err != nil {
			return nil, errors.Wrap(err, "call zclients uic.Auth() error")
		}
		if b {
			auth := new(models.Auth)
			auth.IsOwner = true
			r := make([]*models.NamespaceRole, 0)
			auth.Role = r
			return auth, nil
		}
	}

	//从数据库验证权限
	auth2, err := s.roleService.Find(appId, userId, cluster, env)
	if err != nil {
		return nil, errors.Wrap(err, "call RoleService.Find error")
	}
	auth.Role = auth2.Role
	return auth, nil
}

func (s appService) AuthPerm(userID, perm string) (bool, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	b, err := s.uicService.Auth(ctx, userID, perm, 3)
	if err != nil {
		return false, errors.Wrap(err, "call zclients uic.Auth() error")
	}
	return b, nil
}
