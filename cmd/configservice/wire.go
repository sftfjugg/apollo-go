// +build wireinject

package main

import (
	"github.com/google/wire"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice/controllers"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/app"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/config"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/db"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/httpclient"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/log"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	db.ProviderSet,
	repositories.ProviderSet,
	services.ProviderSet,
	controllers.ProviderSet,
	configservice.ProviderSet,
	http.ProviderSet,
	httpclient.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSets))
}
