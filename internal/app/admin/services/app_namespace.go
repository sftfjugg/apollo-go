package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	_ "go.didapinche.com/time"
	"sort"
	"strconv"
)

type AppNamespaceService interface {
	Create(appNamespace *models.AppNamespace) error
	CreateOrFindAppNamespace(appNamespace *models.AppNamespace) (int64, error)
	DeleteById(id string) error
	DeleteByNameAndAppIdAndCluster(name, appId, cluster string) error
	Update(appNamespace *models.AppNamespace) error
	UpdateIsDisply(appNamespace *models.AppNamespace) error
	FindAllClusterNameByAppId(appId string) ([]string, error)
	FindAppNamespace(appId, cluster, format, comment string) ([]*models2.AppNamespace, error)
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
	FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appId, clusterName, laneName, name string) (*models.AppNamespace, error)
	FindByLaneName(lane string) (*models2.AppPage, error)
}

type appNamespaceService struct {
	db                *gorm.DB
	repository        repositories.AppNamespaceRepository
	itemRepository    repositories.ItemRepisitory
	itemService       ItemService
	releaseRepository repositories.Release
}

func NewAppNamespaceService(
	db *gorm.DB,
	itemRepository repositories.ItemRepisitory,
	itemService ItemService,
	repository repositories.AppNamespaceRepository,
	releaseRepository repositories.Release,
) AppNamespaceService {
	return &appNamespaceService{
		db:                db,
		itemRepository:    itemRepository,
		itemService:       itemService,
		repository:        repository,
		releaseRepository: releaseRepository,
	}
}

func (s appNamespaceService) Create(appNamespace *models.AppNamespace) error {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appNamespace.AppId, appNamespace.ClusterName, appNamespace.LaneName, appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return errors.New("name alrealy exists")
	}
	if appNamespace.IsPublic {
		if appNamespace.Name == "application" || appNamespace.AppId != "public_global_config" {
			return errors.New("公共配置名字不能叫做application，切必须位于公共配置中，暂不支持在其他项目建立公共配置")
		}
	}
	if appNamespace.DeptName == "" {
		dept, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appNamespace.AppId, appNamespace.ClusterName, "default", appNamespace.Name)
		if err != nil {
			return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
		}
		if dept != nil && dept.DeptName != "" {
			appNamespace.DeptName = dept.DeptName
		} else {
			appNamespace.DeptName = "default"
		}
	}
	if appNamespace.Format == "" {
		appNamespace.Format = "服务"
	}

	db := s.db.Begin()
	if err := s.repository.Create(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.Create() error")
	}
	db.Commit()
	return nil
}

//外部rpc调用
func (s appNamespaceService) CreateOrFindAppNamespace(appNamespace *models.AppNamespace) (int64, error) {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appNamespace.AppId, appNamespace.ClusterName, appNamespace.LaneName, appNamespace.Name)
	if err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return int64(app.Id), nil
	}
	if err := s.Create(appNamespace); err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.Create() error")
	}
	createApp, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appNamespace.AppId, appNamespace.ClusterName, appNamespace.LaneName, appNamespace.Name)
	if err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	return int64(createApp.Id), nil
}

func (s appNamespaceService) DeleteById(id string) error {
	app, err := s.repository.FindAppNamespaceById(id)
	if err != nil {
		return errors.Wrap(err, "call repository.FindAppNamespaceById() error")
	}
	db := s.db.Begin()
	if err := s.itemRepository.DeleteByNamespaceId(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.DeleteByIdOnRelease() error")
	}
	if err := s.repository.DeleteById(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.DeleteById() error")
	}
	if err := s.releaseRepository.Delete(db, app.AppId, app.ClusterName, app.Name, app.LaneName); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.Delete() error")
	}
	db.Commit()
	return nil
}

//删除配置文件和对应配置以及发布
func (s appNamespaceService) DeleteByNameAndAppIdAndCluster(name, appId, cluster string) error {
	namespaces, err := s.repository.FindAppNamespaceByAppIdAndName(appId, name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindAppNamespaceByAppIdAndName() error")
	}
	ids := make([]string, 0)
	for _, n := range namespaces {
		ids = append(ids, strconv.FormatUint(n.Id, 10))
	}
	db := s.db.Begin()
	if err := s.itemRepository.DeleteByNamespaceIds(db, ids); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.DeleteByNamespaceIds() error")
	}
	if err := s.repository.DeleteByNameAndAppIdAndCluster(db, name, appId, cluster); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call repository.DeleteByNameAndAppIdAndCluster() error")
	}
	if err := s.releaseRepository.DeleteByName(db, appId, cluster, name); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call releaseRepository.DeleteByName error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) Update(appNamespace *models.AppNamespace) error {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appNamespace.AppId, appNamespace.ClusterName, appNamespace.LaneName, appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" && app.Id != appNamespace.Id {
		return errors.New("name alrealy exists")
	}
	db := s.db.Begin()
	if err := s.repository.Update(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.Update() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) UpdateIsDisply(appNamespace *models.AppNamespace) error {
	db := s.db.Begin()
	if err := s.repository.UpdateIsDisply(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.UpdateIsDisply() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error) {
	appNamespaces, err := s.repository.FindAppNamespaceByAppIdAndClusterName(appId, clusterName)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindAppNamespaceByAppIdAndClusterName() error")
	}
	return appNamespaces, nil
}

func (s appNamespaceService) FindAppNamespace(appId, cluster, format, comment string) ([]*models2.AppNamespace, error) {
	appNamespaces, err := s.repository.FindAppNamespace(appId, cluster, format)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindAppNamespaceByAppId() error")
	}
	if len(appNamespaces) < 1 && format == "" && comment == "" { //如果不存在，则新建一个名为application的
		if cluster == "" {
			cluster = "default"
		}
		appNamespace := new(models.AppNamespace)
		appNamespace.Name = "application"
		appNamespace.ClusterName = cluster
		appNamespace.AppId = appId
		appNamespace.IsPublic = false
		appNamespace.Format = "服务"
		appNamespace.LaneName = "default"
		appNamespace.DeptName = "default"
		appNamespace.IsDisplay = true
		appNamespace.IsDeleted = false
		if err := s.Create(appNamespace); err != nil {
			return nil, errors.Wrap(err, "call appNamespace.Create() error")
		}
		appNamespacess, err := s.repository.FindAppNamespace(appId, cluster, format)
		if err != nil {
			return nil, errors.Wrap(err, "call AppNamespaceRepository.FindAppNamespaceByAppId() error")
		}
		appNamespaces = appNamespacess
	}
	apps := make([]*models2.AppNamespace, 0)

	names := make(map[string][]int)
	for i := range appNamespaces {
		name := appNamespaces[i].AppId + appNamespaces[i].ClusterName + appNamespaces[i].Name
		names[name] = append(names[name], i)
	}

	for k, _ := range names {
		app := new(models2.AppNamespace)
		for _, i := range names[k] {
			namespace := new(models2.Namespace)
			if appNamespaces[i].LaneName == "default" || appNamespaces[i].LaneName == "主版本" {
				app.Format = appNamespaces[i].Format
				app.Name = appNamespaces[i].Name
				app.AppId = appNamespaces[i].AppId
				app.ClusterName = appNamespaces[i].ClusterName
				app.IsPublic = appNamespaces[i].IsPublic
				app.DeptName = appNamespaces[i].DeptName
				app.IsDisplay = appNamespaces[i].IsDisplay
			}
			namespace.Id = appNamespaces[i].Id
			namespace.LaneName = appNamespaces[i].LaneName
			items, err := s.itemService.FindItemByNamespaceId(strconv.FormatUint(namespace.Id, 10), comment)
			if err != nil {
				return nil, errors.Wrap(err, "call itemService.FindItemByNamespaceId() error")
			}
			namespace.Items = items
			app.Namespaces = append(app.Namespaces, namespace)
		}
		apps = append(apps, app)
	}
	sort.Sort(models2.AppNamespaceSlice(apps))
	return apps, nil
}

func (s appNamespaceService) FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appId, clusterName, laneName, name string) (*models.AppNamespace, error) {
	appNamespace, err := s.repository.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appId, clusterName, laneName, name)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane() error")
	}
	return appNamespace, nil
}

func (s appNamespaceService) FindAllClusterNameByAppId(appId string) ([]string, error) {
	appNamespace, err := s.repository.FindAllClusterNameByAppId(appId)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindAllClusterNameByAppId() error")
	}
	names := make([]string, 0)
	for _, n := range appNamespace {
		names = append(names, n.ClusterName)
	}
	return names, nil
}

func (s appNamespaceService) FindByLaneName(lane string) (*models2.AppPage, error) {
	apps, err := s.repository.FindByLaneName(lane)
	if err != nil {
		return nil, errors.Wrap(err, "call repository.FindByLaneName failed")
	}
	appPage := new(models2.AppPage)
	appPage.Total = len(apps)
	appPage.AppNamespaces = apps
	return appPage, nil
}
