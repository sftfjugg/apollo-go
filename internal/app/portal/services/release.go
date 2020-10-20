package services

import (
	"apollo-adminserivce/internal/app/portal/clients"
	models2 "apollo-adminserivce/internal/app/portal/models"
	"github.com/pkg/errors"
	"net/http"
)

type ReleaseService interface {
	Create(env string, r *http.Request) (*models2.Response, error)
}

type releaseService struct {
	httpClient *clients.HttpClient
}

func NewReleaseService(httpClient *clients.HttpClient) ReleaseService {
	return &releaseService{httpClient: httpClient}
}

func (s releaseService) Create(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/release", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}
