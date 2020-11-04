package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
	"sort"
	"strconv"
)

type AppNamespaceService interface {
	Create(appNamespace *models.AppNamespace) error
	CreateOrFindAppNamespace(appNamespace *models.AppNamespace) (int64, error)
	CreateByRelated(appNamespace *models.AppNamespace, items []*models.Item, appName, appId string) error
	DeleteById(id string) error
	DeleteByNameAndAppId(name, appId string) error
	Update(appNamespace *models.AppNamespace) error
	FindAppNamespaceByAppId(appId, format, comment string) ([]*models2.AppNamespace, error)
	FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error)
	FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name string) (*models.AppNamespace, error)
}

type appNamespaceService struct {
	db             *gorm.DB
	repository     repositories.AppNamespaceRepository
	itemRepository repositories.ItemRepisitory
	itemService    ItemService
}

func NewAppNamespaceService(
	db *gorm.DB,
	itemRepository repositories.ItemRepisitory,
	itemService ItemService,
	repository repositories.AppNamespaceRepository,
) AppNamespaceService {
	return &appNamespaceService{
		db:             db,
		itemRepository: itemRepository,
		itemService:    itemService,
		repository:     repository,
	}
}

func (s appNamespaceService) Create(appNamespace *models.AppNamespace) error {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return errors.New("name alrealy exists")
	}
	db := s.db.Begin()
	if err := s.repository.Create(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) CreateOrFindAppNamespace(appNamespace *models.AppNamespace) (int64, error) {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name)
	if err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return int64(app.Id), nil
	}
	if err := s.Create(appNamespace); err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.Create() error")
	}
	createApp, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name)
	if err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	return int64(createApp.Id), nil
}

//通过关联创建
func (s appNamespaceService) CreateByRelated(appNamespace *models.AppNamespace, items []*models.Item, appName, appId string) error {
	if len(items) == 0 {
		return errors.New("The namespace does not have a configuration")
	}
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, "default", appNamespace.Name)
	if err != nil {
		return errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return errors.New("name alrealy exists")
	}
	appNamespace.ClusterName = "default"
	appNamespace.LaneName = "主版本"
	appNamespace.AppId = appId
	appNamespace.IsPublic = true
	id, err := s.FindEmptyAppNamespace()
	if err != nil {
		return errors.Wrap(err, "call AppNamespaceService.FindEmptyAppNamespace() error")
	}
	appNamespace.Id = id
	db := s.db.Begin()
	if err := s.repository.Update(db, appNamespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.Create() error")
	}
	for i := range items {
		items[i].NamespaceId = id
		items[i].DataChange_LastTime = time.Now()
		items[i].DataChange_CreatedTime = time.Now()
	}
	if err := s.itemRepository.Creates(db, items); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Creates() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) DeleteById(id string) error {
	db := s.db.Begin()
	if err := s.itemRepository.DeleteByNamespaceId(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call itemRepository.DeleteByIdOnRelease() error")
	}
	if err := s.repository.DeleteById(db, id); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call AppNamespaceRepository.DeleteById() error")
	}
	db.Commit()
	return nil
}

//删除配置文件和对应配置
func (s appNamespaceService) DeleteByNameAndAppId(name, appId string) error {
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
	if err := s.repository.DeleteByNameAndAppId(db, name, appId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call repository.DeleteByNameAndAppId() error")
	}
	db.Commit()
	return nil
}

func (s appNamespaceService) Update(appNamespace *models.AppNamespace) error {
	app, err := s.FindOneAppNamespaceByAppIdAndClusterNameAndName(appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name)
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

func (s appNamespaceService) FindAppNamespaceByAppIdAndClusterName(appId, clusterName string) ([]*models.AppNamespace, error) {
	appNamespaces, err := s.repository.FindAppNamespaceByAppIdAndClusterName(appId, clusterName)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindAppNamespaceByAppIdAndClusterName() error")
	}
	return appNamespaces, nil
}

func (s appNamespaceService) FindAppNamespaceByAppId(appId, format, comment string) ([]*models2.AppNamespace, error) {
	appNamespaces, err := s.repository.FindAppNamespaceByAppId(appId, format)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindAppNamespaceByAppId() error")
	}
	apps := make([]*models2.AppNamespace, 0)
	names := make(map[string][]int)
	for i, a := range appNamespaces {
		names[a.Name] = append(names[a.Name], i)
	}
	for k := range names {
		app := new(models2.AppNamespace)
		app.Name = k
		for i := range names[k] {
			j := names[k][i]
			namespace := new(models2.Namespace)
			app.AppId = appNamespaces[j].AppId
			app.IsPublic = appNamespaces[j].IsPublic
			if appNamespaces[j].Format != "" {
				app.Format = appNamespaces[j].Format
			}
			namespace.ClusterName = appNamespaces[j].ClusterName
			namespace.Id = appNamespaces[j].Id
			namespace.LaneName = appNamespaces[j].LaneName
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
func (s appNamespaceService) FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name string) (*models.AppNamespace, error) {
	appNamespace, err := s.repository.FindOneAppNamespaceByAppIdAndClusterNameAndName(appId, clusterName, name)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRepository.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	return appNamespace, nil
}

//占坑，如果存在Idonotexist返回id，不存在创建一个没有用的AppNamespace，返回id
func (s appNamespaceService) FindEmptyAppNamespace() (uint64, error) {
	appNamespace, err := s.repository.FindOneAppNamespaceByAppIdAndClusterNameAndName("Idonotexist", "Idonotexist", "Idonotexist")
	if err != nil {
		return 0, errors.Wrap(err, "call AppNamespaceRepository.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if appNamespace.Id == 0 {
		app := new(models.AppNamespace)
		app.Name = "Idonotexist"
		app.ClusterName = "Idonotexist"
		app.AppId = "Idonotexist"
		s.Create(app)
		app, err := s.repository.FindOneAppNamespaceByAppIdAndClusterNameAndName("Idonotexist", "Idonotexist", "Idonotexist")
		if err != nil {
			return 0, errors.Wrap(err, "call AppNamespaceRepository.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
		}
		appNamespace.Id = app.Id
	}
	return appNamespace.Id, nil
}
