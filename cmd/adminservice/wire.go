package main

import (
	"apollo-adminserivce/internal/app/adminservice"
	"apollo-adminserivce/internal/app/adminservice/controllers"
	"apollo-adminserivce/internal/app/adminservice/repositories"
	"apollo-adminserivce/internal/app/adminservice/services"
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
