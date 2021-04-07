// +build wireinject

package main

import (
	"github.com/google/wire"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/address"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/controllers"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zservice"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/app"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/config"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/db"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/dingding"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/httpclient"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/log"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/zeus"
	"go.didapinche.com/foundation/ophis"
	"go.didapinche.com/uic"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	db.ProviderSet,
	zeus.ProviderSet,
	zclients.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	controllers.ProviderSet,
	address.ProviderSet,
	http.ProviderSet,
	httpclient.ProviderSet,
	portal.ProviderSet,
	uic.ProviderSet,
	zservice.ProviderSet,
	dingding.ProviderSet,
	ophis.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
