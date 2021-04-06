package services

import (
	"github.com/pkg/errors"
	"github.com/uber/tchannel-go"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/repositories"
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
	FindAuths(userId, cluster, env string) (map[string]*models.Auth, error)
	GetFavorite(userId string) ([]*plat_limos_rpc.App, error) //获得收藏的limos项目
	GetOwner(userId string) ([]*plat_limos_rpc.App, error)    //获得负责的limos项目
	GetRecent(userId string) ([]*plat_limos_rpc.App, error)   //获得最近的limos项目
}

type appService struct {
	limosService plat_limos_rpc.TChanLimosService
	uicService   uic_service_api.TChanUicService
	roleService  RoleService
	history      repositories.HistoryRepository
}

func NewAppService(
	limosService plat_limos_rpc.TChanLimosService,
	uicService uic_service_api.TChanUicService,
	history repositories.HistoryRepository,
	roleService RoleService,
) AppService {
	return &appService{
		limosService: limosService,
		uicService:   uicService,
		history:      history,
		roleService:  roleService,
	}
}

func (s appService) GetOwner(userId string) ([]*plat_limos_rpc.App, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	apps, err := s.limosService.FindAppForPage(ctx, "", "", "", "", userId, 0, "mine", 100, 1, "", 0)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.FindAppForPage() error")
	}
	return apps.Apps, nil
}

func (s appService) GetFavorite(userId string) ([]*plat_limos_rpc.App, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	apps, err := s.limosService.FindAppForPage(ctx, "", "", "", "", userId, 0, "favorite", 100, 1, "", 0)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.FindAppForPage() error")
	}
	return apps.Apps, nil
}

func (s appService) GetRecent(userId string) ([]*plat_limos_rpc.App, error) {
	limosApps := make([]*plat_limos_rpc.App, 0)
	apps, err := s.history.FindHistory(userId)
	if err != nil {
		return nil, errors.Wrap(err, "call history FindHistory() error")
	}
	for _, app := range apps {
		limosApp := new(plat_limos_rpc.App)
		limosApp.Name = app.AppId
	}
	return limosApps, nil
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

func (s appService) FindAuths(userId, cluster, env string) (map[string]*models.Auth, error) {
	auths := make(map[string]*models.Auth)
	//验证root权限
	b, err := s.AuthPerm(userId, constans.ApolloRoot)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.Auth() error")
	}
	if b {
		auth := new(models.Auth)
		auth.IsRoot = true
		auth.IsOperate = true
		auth.IsOwner = true
		r := make([]*models.NamespaceRole, 0)
		auth.Role = r
		auths["root"] = auth
	}

	//公共配置权限默认为owner，不受op管理
	isPublic, err := s.AuthPerm(userId, constans.ApolloPublicOperate)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.Auth() error")
	}
	if isPublic {
		auth := new(models.Auth)
		auth.IsRoot = true
		r := make([]*models.NamespaceRole, 0)
		auth.Role = r
		auth.IsOperate = true
		auth.IsOwner = true
		auths["public_global_config"] = auth
	}

	isOp, err := s.AuthPerm(userId, constans.AppOperate)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.Auth() error")
	}
	if isOp {
		auth := new(models.Auth)
		auth.IsRoot = false
		r := make([]*models.NamespaceRole, 0)
		auth.Role = r
		auth.IsOperate = true
		auth.IsOwner = true
		auths["operate"] = auth
	}

	//负责的项目
	apps, err := s.GetOwner(userId)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.GetAppByID() error")
	}
	for _, app := range apps {
		auth := new(models.Auth)
		auth.IsRoot = false
		r := make([]*models.NamespaceRole, 0)
		auth.Role = r
		auth.IsOperate = false
		auth.IsOwner = true
		auths[app.Name] = auth
	}

	//从数据库验证权限
	auth2, err := s.roleService.Finds(userId, cluster, env)
	if err != nil {
		return nil, errors.Wrap(err, "call RoleService.Find error")
	}
	for k, v := range auth2 {
		if k1, ok := auths[k]; ok {
			k1.Role = v
		} else {
			auth := new(models.Auth)
			auth.IsRoot = false
			auth.Role = v
			auth.IsOperate = false
			auth.IsOwner = false
			auths[k] = auth
		}
	}
	return auths, nil
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

	//验证owner权限和op权限，公共配置直接为root权限
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
			auth.IsRoot = true
			r := make([]*models.NamespaceRole, 0)
			auth.Role = r
			return auth, nil
		}
	}
	if auth.IsOperate && auth.IsOwner {
		auth.IsRoot = true
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
