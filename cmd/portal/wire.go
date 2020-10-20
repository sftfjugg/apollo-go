package main

import (
	"apollo-adminserivce/internal/app/portal"
	"apollo-adminserivce/internal/app/portal/address"
	"apollo-adminserivce/internal/app/portal/clients"
	"apollo-adminserivce/internal/app/portal/controllers"
	"apollo-adminserivce/internal/app/portal/repositories"
	"apollo-adminserivce/internal/app/portal/services"
	"apollo-adminserivce/internal/pkg/app"
	"apollo-adminserivce/internal/pkg/config"
	"apollo-adminserivce/internal/pkg/db"
	"apollo-adminserivce/internal/pkg/http"
	"apollo-adminserivce/internal/pkg/httpclient"
	"apollo-adminserivce/internal/pkg/log"
	"apollo-adminserivce/internal/pkg/zeus"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	db.ProviderSet,
	zeus.ProviderSet,
	clients.ProviderSet,
	repositories.ProviderSet,
	services.ProviderSet,
	controllers.ProviderSet,
	address.ProviderSet,
	http.ProviderSet,
	httpclient.ProviderSet,
	portal.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSets))
}
