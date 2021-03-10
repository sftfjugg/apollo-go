package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/uber/tchannel-go"
	models22 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"go.didapinche.com/goapi/plat_limos_rpc"
	"net/http"
	"time"
)

type AppNamespaceService interface {
	Create(env string, r *http.Request) (*models2.Response, error)
	CreateLane(env string, r *http.Request) (*models2.Response, error)
	CreateCluster(env string, r *http.Request) (*models2.Response, error)
	DeleteById(env string, r *http.Request) (*models2.Response, error)
	DeleteByNameAndAppId(env string, r *http.Request) (*models2.Response, error)
	Update(env string, r *http.Request) (*models2.Response, error)
	UpdateIsDisply(env string, r *http.Request) (*models2.Response, error)
	FindAllClusterNameByAppId(r *http.Request) (*models2.Response, error)
	FindAppNamespaceByAppId(appId string, r *http.Request) (*models2.Response, error)
	FindByLaneName(r *http.Request) (*models2.Response, error)
	FindAppNamespaceByAppIdAndClusterName(env string, r *http.Request) (*models2.Response, error)
	FindAppByLaneNameandAppId(r *http.Request) (*models2.Response, error)
}

type appNamespaceService struct {
	limosService plat_limos_rpc.TChanLimosService
	httpClient   *zclients.HttpClient
}

func NewAppNamespaceService(
	limosService plat_limos_rpc.TChanLimosService,
	httpClient *zclients.HttpClient,
) AppNamespaceService {
	return appNamespaceService{
		httpClient:   httpClient,
		limosService: limosService,
	}
}

func (s appNamespaceService) Create(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

func (s appNamespaceService) CreateLane(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/app_namespace_by_lane", env, r)
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

//查询某泳道在所有环境下的数目
func (s appNamespaceService) FindByLaneName(r *http.Request) (*models2.Response, error) {
	param := new(struct {
		Test   *models22.AppPage `json:"test"`
		Aliyun *models22.AppPage `json:"aliyun"`
		Online *models22.AppPage `json:"online"`
		Total  int               `json:"total"`
	})
	total := 0
	response, err := s.httpClient.HttpDo("/app_by_lane", "TEST", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response.Code == 200 {
		test := new(models22.AppPage)
		if err := json.Unmarshal(response.Data, &test); err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal falied")
		}
		param.Test = test
		total += test.Total
	}

	response2, err := s.httpClient.HttpDo("/app_by_lane", "ONLINE", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response2.Code == 200 {
		online := new(models22.AppPage)
		if err := json.Unmarshal(response2.Data, &online); err != nil {

		}
		param.Online = online
		total += online.Total
	}

	response3, err := s.httpClient.HttpDo("/app_by_lane", "ALIYUN", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response3.Code == 200 {
		aliyun := new(models22.AppPage)
		if err := json.Unmarshal(response3.Data, &aliyun); err != nil {

		}
		param.Aliyun = aliyun
		total += aliyun.Total
	}
	param.Total = total
	response.Data, _ = json.Marshal(param)
	return response, nil
}

func (s appNamespaceService) FindAllClusterNameByAppId(r *http.Request) (*models2.Response, error) {

	param := new(struct {
		TEST   []string `json:"test"`
		ALIYUN []string `json:"aliyun"`
		ONLINE []string `json:"online"`
	})
	tests := make([]string, 0)
	response, err := s.httpClient.HttpDo("/cluster", "TEST", r)
	if err != nil {
		tests = append(tests, "default")
		//return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	} else {
		if err := json.Unmarshal(response.Data, &tests); err != nil {
			tests = append(tests, "default")
		}
		for i, v := range tests {
			if v == "default" {
				tests[i] = tests[0]
				tests[0] = v
			}
		}
	}
	param.TEST = tests
	aliyuns := make([]string, 0)
	response, err = s.httpClient.HttpDo("/cluster", "ALIYUN", r)
	if err != nil {
		aliyuns = append(aliyuns, "default")
		//return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	} else {
		if err := json.Unmarshal(response.Data, &aliyuns); err != nil {
			aliyuns = append(aliyuns, "default")
		}
		for i, v := range aliyuns {
			if v == "default" {
				aliyuns[i] = aliyuns[0]
				aliyuns[0] = v
			}
		}
	}
	param.ALIYUN = aliyuns
	online := make([]string, 0)
	response, err = s.httpClient.HttpDo("/cluster", "ONLINE", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	} else {
		if err := json.Unmarshal(response.Data, &online); err != nil {
			online = append(online, "default")
		}
		for i, v := range online {
			if v == "default" {
				online[i] = online[0]
				online[0] = v
			}
		}
	}
	param.ONLINE = online
	response.Code = 200
	response.Data, _ = json.Marshal(param)
	response.ContentType = "application/json; charset=utf-8"
	return response, nil
}

//查询应用在某泳道下所有环境下的数目
func (s appNamespaceService) FindAppByLaneNameandAppId(r *http.Request) (*models2.Response, error) {
	param := new(struct {
		Test   *models22.AppPage `json:"test"`
		Aliyun *models22.AppPage `json:"aliyun"`
		Online *models22.AppPage `json:"online"`
		Total  int               `json:"total"`
	})
	total := 0
	response, err := s.httpClient.HttpDo("/app_by_app_and_lane", "TEST", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response.Code == 200 {
		test := new(models22.AppPage)
		if err := json.Unmarshal(response.Data, &test); err != nil {

		}
		for i, _ := range test.AppNamespaces {
			limos, err := s.FindLimosApp(test.AppNamespaces[i].AppId)
			if err == nil {
				test.AppNamespaces[i].Owner = limos.Owner
				test.AppNamespaces[i].Level = limos.Level
				test.AppNamespaces[i].LimosId = limos.ID
			}
		}
		param.Test = test
		total += test.Total
	}

	response2, err := s.httpClient.HttpDo("/app_by_app_and_lane", "ONLINE", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response2.Code == 200 {
		online := new(models22.AppPage)
		if err := json.Unmarshal(response2.Data, &online); err != nil {

		}
		for i, _ := range online.AppNamespaces {
			limos, err := s.FindLimosApp(online.AppNamespaces[i].AppId)
			if err == nil {
				online.AppNamespaces[i].Owner = limos.Owner
				online.AppNamespaces[i].Level = limos.Level
				online.AppNamespaces[i].LimosId = limos.ID
			}
		}
		param.Online = online
		total += online.Total
	}

	response3, err := s.httpClient.HttpDo("/app_by_app_and_lane", "ALIYUN", r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response3.Code == 200 {
		aliyun := new(models22.AppPage)
		if err := json.Unmarshal(response3.Data, &aliyun); err != nil {

		}
		param.Aliyun = aliyun
		for i, _ := range aliyun.AppNamespaces {
			limos, err := s.FindLimosApp(aliyun.AppNamespaces[i].AppId)
			if err == nil {
				aliyun.AppNamespaces[i].Owner = limos.Owner
				aliyun.AppNamespaces[i].Level = limos.Level
				aliyun.AppNamespaces[i].LimosId = limos.ID
			}
		}
		total += aliyun.Total
	}
	param.Total = total
	response.Data, _ = json.Marshal(param)
	return response, nil
}

func (s appNamespaceService) FindLimosApp(name string) (*plat_limos_rpc.App, error) {
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	apps, err := s.limosService.FindAppForPage(ctx, name, "", "", "", "", 0, "all", 0, 20, "", 0)
	if err != nil {
		return nil, errors.Wrap(err, "call zclients uic.FindAppForPage() error")
	}
	app := new(plat_limos_rpc.App)
	for _, v := range apps.Apps {
		if v.Name == name {
			app = v
			break
		}
	}
	return app, nil
}
