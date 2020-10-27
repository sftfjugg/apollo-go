package services

import (
	"apollo-adminserivce/internal/app/portal/clients"
	models2 "apollo-adminserivce/internal/app/portal/models"
	"github.com/pkg/errors"
	"net/http"
)

type ItemService interface {
	Create(env string, r *http.Request) (*models2.Response, error)
	Update(env string, r *http.Request) (*models2.Response, error)
	DeleteByNamespaceIdAndKey(env string, r *http.Request) (*models2.Response, error)
	DeleteByNamespaceId(env string, r *http.Request) (*models2.Response, error)
	FindItemByNamespaceId(env string, r *http.Request) (*models2.Response, error)
	FindItemByNamespaceIdOnRelease(env string, r *http.Request) (*models2.Response, error)
	FindItemByAppIdAndKey(env string, r *http.Request) (*models2.Response, error)
	FindItemByKeyForPage(env string, r *http.Request) (*models2.Response, error)
	FindAppItemByKeyForPage(env string, r *http.Request) (*models2.Response, error)
	FindItemByNamespaceIdAndKey(env string, r *http.Request) (*models2.Response, error)
}

type itemService struct {
	httpClient *clients.HttpClient
}

func NewItemService(httpClient *clients.HttpClient) ItemService {
	return &itemService{httpClient: httpClient}
}

func (s itemService) Create(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/item", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) Update(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/item", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) DeleteByNamespaceIdAndKey(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/item", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) DeleteByNamespaceId(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/items", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) FindItemByNamespaceId(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/items", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) FindItemByNamespaceIdOnRelease(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/items_release", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) FindItemByAppIdAndKey(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/item_by_key_and_appId", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) FindItemByKeyForPage(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/items_by_key", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) FindAppItemByKeyForPage(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/items_by_key_on_app", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s itemService) FindItemByNamespaceIdAndKey(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/item", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}
