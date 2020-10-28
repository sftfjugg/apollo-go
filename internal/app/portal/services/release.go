package services

import (
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"net/http"
)

type ReleaseService interface {
	Create(env string, r *http.Request) (*models2.Response, error)
}

type releaseService struct {
	httpClient *zclients.HttpClient
}

func NewReleaseService(httpClient *zclients.HttpClient) ReleaseService {
	return &releaseService{httpClient: httpClient}
}

func (s releaseService) Create(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/release", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}
