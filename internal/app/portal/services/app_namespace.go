package services

import (
	"apollo-adminserivce/internal/app/portal/clients"
	models2 "apollo-adminserivce/internal/app/portal/models"
	"github.com/pkg/errors"
	"net/http"
)

type AppNamespaceService interface {
	Create(env string, r *http.Request) (*models2.Response, error)
	DeleteById(env string, r *http.Request) (*models2.Response, error)
	Update(env string, r *http.Request) (*models2.Response, error)
	FindAppNamespaceByAppIdAndClusterName(env string, r *http.Request) (*models2.Response, error)
}

type appNamespaceService struct {
	httpClient *clients.HttpClient
}

func NewAppNamespaceService(httpClient *clients.HttpClient) AppNamespaceService {
	return appNamespaceService{httpClient: httpClient}
}

func (s appNamespaceService) Create(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace", env, r)
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
