package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"net/http"
)

type AppNamespaceService interface {
	Create(env string, r *http.Request) (*models2.Response, error)
	CreateCluster(env string, r *http.Request) (*models2.Response, error)
	DeleteById(env string, r *http.Request) (*models2.Response, error)
	DeleteByNameAndAppId(env string, r *http.Request) (*models2.Response, error)
	Update(env string, r *http.Request) (*models2.Response, error)
	UpdateIsDisply(env string, r *http.Request) (*models2.Response, error)
	FindAllClusterNameByAppId(r *http.Request) (*models2.Response, error)
	FindAppNamespaceByAppId(appId string, r *http.Request) (*models2.Response, error)
	FindByLaneName(env string, r *http.Request) (*models2.Response, error)
	FindAppNamespaceByAppIdAndClusterName(env string, r *http.Request) (*models2.Response, error)
}

type appNamespaceService struct {
	httpClient *zclients.HttpClient
}

func NewAppNamespaceService(
	httpClient *zclients.HttpClient,
) AppNamespaceService {
	return appNamespaceService{
		httpClient: httpClient,
	}
}

func (s appNamespaceService) Create(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s appNamespaceService) CreateCluster(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/cluster", env, r)
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

func (s appNamespaceService) DeleteByNameAndAppId(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace_by_name", env, r)
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

func (s appNamespaceService) UpdateIsDisply(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace_is_dispaly", env, r)
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

func (s appNamespaceService) FindByLaneName(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_by_lane", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s appNamespaceService) FindAllClusterNameByAppId(r *http.Request) (*models2.Response, error) {

	param := new(struct {
		TEST   []string `json:"test"`
		ALIYUN []string `json:"aliyun"`
		ONLINE []string `json:"online"`
	})
	response, err := s.httpClient.HttpDo("/cluster", "TEST", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	tests := make([]string, 0)
	if err := json.Unmarshal(response.Data, &tests); err != nil {
		tests = append(tests, "default")
	}
	for i, v := range tests {
		if v == "default" {
			tests[i] = tests[0]
			tests[0] = v
		}
	}
	param.TEST = tests
	response, err = s.httpClient.HttpDo("/cluster", "ALIYUN", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	aliyuns := make([]string, 0)
	if err := json.Unmarshal(response.Data, &aliyuns); err != nil {
		aliyuns = append(aliyuns, "default")
	}
	for i, v := range aliyuns {
		if v == "default" {
			aliyuns[i] = aliyuns[0]
			aliyuns[0] = v
		}
	}
	param.ALIYUN = aliyuns
	response, err = s.httpClient.HttpDo("/cluster", "ONLINE", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	online := make([]string, 0)
	if err := json.Unmarshal(response.Data, &online); err != nil {
		online = append(online, "default")
	}
	for i, v := range online {
		if v == "default" {
			online[i] = online[0]
			online[0] = v
		}
	}
	param.ONLINE = online
	response.Code = 200
	response.Data, _ = json.Marshal(param)
	response.ContentType = "application/json; charset=utf-8"
	return response, nil
}
