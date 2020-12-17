package services

import (
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/goapi/apollo_thrift_service/v2"
	"strconv"
)

type ZserviceApi interface {
	CreateOrFindAppNamespace(app *apollo_thrift_service.AppNamespace) (int64, error)
	CreateOrUpdateItem(item *apollo_thrift_service.Item) error
	PublishNamespace(release *apollo_thrift_service.Release) error
}

type zserviceApi struct {
	httpClient *zclients.HttpClient
}

func NewZserviceApi(httpClient *zclients.HttpClient) ZserviceApi {
	return &zserviceApi{httpClient: httpClient}
}

func (s zserviceApi) CreateOrFindAppNamespace(app *apollo_thrift_service.AppNamespace) (int64, error) {

	appNamespace := new(models.AppNamespace)
	appNamespace.Name = app.Name
	appNamespace.AppId = app.AppId
	appNamespace.Format = app.Format
	appNamespace.Comment = app.Comment
	appNamespace.LaneName = app.LaneName
	appNamespace.ClusterName = app.ClusterName
	appNamespace.DataChange_LastModifiedBy = app.Operator
	appNamespace.DataChange_CreatedBy = app.Operator
	env := s.EnvToString(app.Env)
	resp, err := s.httpClient.HttpPost("/app_namespace/create_or_find", env, appNamespace)
	if err != nil {
		return 0, errors.Wrap(err, "HttpClient HttpPost run failed")
	}
	if resp.Code != 200 {
		return 0, errors.New(string(resp.Data))
	}
	id, err := strconv.ParseInt(string(resp.Data), 10, 64)
	return id, nil
}

func (s zserviceApi) CreateOrUpdateItem(item *apollo_thrift_service.Item) error {
	item2 := new(models.Item)
	item2.NamespaceId = uint64(item.NamespaceId)
	item2.Key = item.Key
	item2.Value = item.Value
	item2.DataChange_LastModifiedBy = item.Operator
	env := s.EnvToString(item.Env)
	resp, err := s.httpClient.HttpPost("/item/create_or_update", env, item2)
	if err != nil {
		return errors.Wrap(err, "HttpClient HttpPost run failed")
	}
	if resp.Code != 200 {
		return errors.New(string(resp.Data))
	}
	return nil
}

func (s zserviceApi) PublishNamespace(release *apollo_thrift_service.Release) error {
	param := new(struct {
		Comment     string   `json:"comment"`
		NamespaceId uint64   `json:"namespace_id"`
		Keys        []string `json:"keys"`
		Operator    string   `json:"operator"`
	})
	param.NamespaceId = uint64(release.NamespaceId)
	param.Comment = release.Comment
	param.Keys = release.Keys
	param.Operator = release.Operator
	env := s.EnvToString(release.Env)
	resp, err := s.httpClient.HttpPost("/publish_namespace", env, param)
	if err != nil {
		return errors.Wrap(err, "HttpClient HttpPost run failed")
	}
	if resp.Code != 200 {
		return errors.New(string(resp.Data))
	}
	return nil
}

//env杯举转字符串
func (s zserviceApi) EnvToString(env string) string {
	if env == "1" || env == "TEST" {
		return "TEST"
	} else if env == "4" || env == "ALIYUN" {
		return "ALIYUN"
	} else if env == "3" || env == "ONLINE" {
		return "ONLINE"
	} else {
		return "TEST"
	}
}
