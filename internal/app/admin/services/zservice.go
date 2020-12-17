package services

import (
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type ZService interface {
	CreateOrFindAppNamespace(appNamespace *models.AppNamespace) (int64, error)
	CreateOrUpdateItem(item *models.Item) error
	PublishNamespace(appId, clusterName, comment, name, namespaceId, laneName, operator string, keys []string) error
}

type zService struct {
	appNamespaceService   AppNamespaceService
	itemService           ItemService
	releaseMessageService ReleaseMessageService
}

func NewZService(
	appNamespaceService AppNamespaceService,
	itemService ItemService,
	releaseMessageService ReleaseMessageService,
) ZService {
	return &zService{
		appNamespaceService:   appNamespaceService,
		itemService:           itemService,
		releaseMessageService: releaseMessageService,
	}
}

//外部rpc调用
func (s zService) CreateOrFindAppNamespace(appNamespace *models.AppNamespace) (int64, error) {
	app, err := s.appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appNamespace.AppId, appNamespace.ClusterName, appNamespace.LaneName, appNamespace.Name)
	if err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	if app.Name != "" {
		return int64(app.Id), nil
	}
	if err := s.appNamespaceService.Create(appNamespace); err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.Create() error")
	}
	createApp, err := s.appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndNameAndLane(appNamespace.AppId, appNamespace.ClusterName, appNamespace.LaneName, appNamespace.Name)
	if err != nil {
		return 0, errors.Wrap(err, "call appNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName() error")
	}
	return int64(createApp.Id), nil
}

func (s zService) CreateOrUpdateItem(item *models.Item) error {
	item2, err := s.itemService.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	item.Id = item2.Id
	if item2.Key != "" {
		if err := s.itemService.Update(item); err != nil {
			return errors.Wrap(err, "call itemService.Update() error")
		}
	} else {
		item.DataChange_CreatedBy = item.DataChange_LastModifiedBy
		if err := s.itemService.Create(item); err != nil {
			return errors.Wrap(err, "call itemService.Create() error")
		}
	}
	return nil
}

func (s zService) PublishNamespace(appId, clusterName, comment, name, namespaceId, laneName, operator string, keys []string) error {
	if err := s.releaseMessageService.Create(appId, clusterName, comment, name, namespaceId, laneName, operator, keys); err != nil {
		return err
	}
	return nil
}
