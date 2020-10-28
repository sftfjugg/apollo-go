package zservice

import (
	"github.com/google/wire"
	"github.com/uber/tchannel-go/thrift"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"go.didapinche.com/goapi/apollo_thrift_service/v2"
)

type ZserviceApi struct {
	service services.ZserviceApi
}

func NewZserviceApi(service services.ZserviceApi) apollo_thrift_service.TChanApolloThriftService {
	return &ZserviceApi{service: service}
}

func (z ZserviceApi) CreateOrFindAppNamespace(ctx thrift.Context, app *apollo_thrift_service.AppNamespace) (int64, error) {
	return z.service.CreateOrFindAppNamespace(app)
}

func (z ZserviceApi) CreateOrUpdateItem(ctx thrift.Context, item *apollo_thrift_service.Item) error {
	return z.service.CreateOrUpdateItem(item)
}

func (z ZserviceApi) PublicNamespace(ctx thrift.Context, release *apollo_thrift_service.Release) error {
	return z.service.PublicNamespace(release)
}

var ProviderSet = wire.NewSet(
	NewZserviceApi,
)
