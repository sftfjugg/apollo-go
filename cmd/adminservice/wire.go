package main

import (
	"github.com/google/wire"
	"go.didapinche.com/foundation/apollo-plus/internal/app/adminservice"
	"go.didapinche.com/foundation/apollo-plus/internal/app/adminservice/controllers"
	"go.didapinche.com/foundation/apollo-plus/internal/app/adminservice/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/app/adminservice/services"
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
	controllers.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	adminservice.ProviderSet,
	http.ProviderSet,
	httpclient.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSets))
}
