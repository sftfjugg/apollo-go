// +build wireinject

package main

import (
	"apollo-adminserivce/internal/app/configservice"
	"apollo-adminserivce/internal/app/configservice/controllers"
	"apollo-adminserivce/internal/app/configservice/repositories"
	"apollo-adminserivce/internal/app/configservice/services"
	"apollo-adminserivce/internal/pkg/app"
	"apollo-adminserivce/internal/pkg/config"
	"apollo-adminserivce/internal/pkg/db"
	"apollo-adminserivce/internal/pkg/http"
	"apollo-adminserivce/internal/pkg/httpclient"
	"apollo-adminserivce/internal/pkg/log"
	"github.com/google/wire"
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
