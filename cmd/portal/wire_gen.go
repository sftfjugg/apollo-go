// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

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
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/httpclient"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/log"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/zeus"
	"go.didapinche.com/uic"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	portalOptions, err := portal.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	zeusZeus, err := zeus.New(viper)
	if err != nil {
		return nil, err
	}
	tChanLimosService, err := zclients.NewLimosService(zeusZeus)
	if err != nil {
		return nil, err
	}
	tChanUicService, err := zclients.NewUicService(zeusZeus)
	if err != nil {
		return nil, err
	}
	dbOptions, err := db.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	gormDB, err := db.New(dbOptions, logger)
	if err != nil {
		return nil, err
	}
	roleRepository := repositories.NewRoleReposotory(gormDB)
	roleService := services.NewRoleService(roleRepository, gormDB)
	appService := services.NewAppService(tChanLimosService, tChanUicService, roleService)
	appController := controllers.NewAppController(appService)
	uicOptions := uic.NewOptions(viper)
	api, err := uic.NewApi(uicOptions, logger, tChanUicService)
	if err != nil {
		return nil, err
	}
	client := httpclient.New()
	httpClient := zclients.NewHttpClient(client)
	appNamespaceService := services.NewAppNamespaceService(httpClient)
	appNamespaceController := controllers.NewAppNamespaceController(appNamespaceService)
	itemService := services.NewItemService(httpClient)
	itemController := controllers.NewItemController(itemService)
	releaseService := services.NewReleaseService(httpClient)
	releaseController := controllers.NewReleaseController(releaseService)
	roleController := controllers.NewRoleController(roleService)
	releaseHistoryService := services.NewReleaseHistoryService(httpClient)
	releaseHistoryController := controllers.NewReleaseHistoryController(releaseHistoryService)
	initControllers := controllers.InitControllersFn(appController, api, appNamespaceController, itemController, releaseController, roleController, releaseHistoryController)
	engine, err := http.NewRouter(httpOptions, logger, initControllers)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	zserviceApi := services.NewZserviceApi(httpClient)
	tChanApolloThriftService := zservice.NewZserviceApi(zserviceApi)
	serverServer := zeus.NewZeusServer(zeusZeus, tChanApolloThriftService)
	addressOptions, err := address.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	meta, err := address.NewMetas(addressOptions)
	if err != nil {
		return nil, err
	}
	addressService := services.NewAddress(meta, logger)
	application, err := portal.NewApp(portalOptions, logger, server, serverServer, addressService)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, db.ProviderSet, zeus.ProviderSet, zclients.ProviderSet, services.ProviderSet, repositories.ProviderSet, controllers.ProviderSet, address.ProviderSet, http.ProviderSet, httpclient.ProviderSet, portal.ProviderSet, uic.ProviderSet, zservice.ProviderSet)
