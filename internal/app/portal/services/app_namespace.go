package services

import (
	"apollo-adminserivce/internal/app/portal/clients"
	models2 "apollo-adminserivce/internal/app/portal/models"
	"apollo-adminserivce/internal/app/portal/repositories"
	"github.com/pkg/errors"
	"net/http"
)

type AppNamespaceService interface {
	Create(env string, r *http.Request) (*models2.Response, error)
	DeleteById(env string, r *http.Request) (*models2.Response, error)
	Update(env string, r *http.Request) (*models2.Response, error)
	CreateByRelated(namespaceId, appName, appId, env string) (*models2.Response, error)
	FindAppNamespaceByAppId(env string, r *http.Request) (*models2.Response, error)
	FindAppNamespaceByAppIdAndClusterName(env string, r *http.Request) (*models2.Response, error)
}

type appNamespaceService struct {
	httpClient     *clients.HttpClient
	repository     repositories.AppNamespaceRelatedRepository
	itemRepository repositories.ItemRelatedRepisitory
}

func NewAppNamespaceService(
	httpClient *clients.HttpClient,
	repository repositories.AppNamespaceRelatedRepository,
	itemRepository repositories.ItemRelatedRepisitory,
) AppNamespaceService {
	return appNamespaceService{
		httpClient:     httpClient,
		repository:     repository,
		itemRepository: itemRepository,
	}
}

func (s appNamespaceService) Create(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s appNamespaceService) CreateByRelated(namespaceId, appName, appId, env string) (*models2.Response, error) {
	appNamespace, err := s.repository.FindAppNamespaceById(namespaceId)
	if err != nil {
		return nil, errors.Wrap(err, "call AppNamespaceRelatedRepository.FindAppNamespaceById error")
	}
	items, err := s.itemRepository.FindItemByNamespaceId(namespaceId)
	if err != nil {
		return nil, errors.Wrap(err, "call itemRepository.FindItemByNamespaceId error")
	}
	param := new(struct {
		AppNamespace *models2.AppNamespace `json:"app_namespace"`
		Items        []*models2.Item       `json:"items"`
		AppName      string                `json:"app_name"`
		AppId        string                `json:"app_id"`
	})
	param.AppNamespace = appNamespace
	param.Items = items
	param.AppName = appName
	param.AppId = appId
	response, err := s.httpClient.HttpPost("/app_namespace_related", env, param)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s appNamespaceService) DeleteById(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s appNamespaceService) Update(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s appNamespaceService) FindAppNamespaceByAppIdAndClusterName(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}
func (s appNamespaceService) FindAppNamespaceByAppId(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace_all", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}
