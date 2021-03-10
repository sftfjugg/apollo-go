package services

import (
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"net/http"
)

type ReleaseHistoryService interface {
	Find(env string, r *http.Request) (*models2.Response, error)
}

type releaseHistoryService struct {
	httpClient *zclients.HttpClient
}

func NewReleaseHistoryService(httpClient *zclients.HttpClient) ReleaseHistoryService {
	return &releaseHistoryService{httpClient: httpClient}
}

func (s releaseHistoryService) Find(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/release_history", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}
