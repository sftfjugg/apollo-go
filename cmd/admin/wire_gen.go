// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/controllers"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/app"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/config"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/db"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/httpclient"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/log"
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
	adminOptions, err := admin.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
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
	itemRepisitory := repositories.NewItemRepisitory(gormDB)
	itemService := services.NewItemService(itemRepisitory, gormDB)
	appNamespaceRepository := repositories.NewAppNamespaceRepository(gormDB)
	appNamespaceService := services.NewAppNamespaceService(gormDB, itemRepisitory, itemService, appNamespaceRepository)
	appNamespaceController := controllers.NewAppNamespaceController(appNamespaceService)
	itemController := controllers.NewItemController(itemService)
	releaseHistoryRepository := repositories.NewReleaseHistoryRepository(gormDB)
	releaseHistoryService := services.NewReleaseHistoryService(gormDB, releaseHistoryRepository)
	releaseHistoryController := controllers.NewReleaseHistoryController(releaseHistoryService)
	releaseMessageRepository := repositories.NewReleaseMessageRepository()
	release := repositories.NewRelease(gormDB)
	releaseMessageService := services.NewReleaseMessageService(releaseMessageRepository, release, appNamespaceRepository, releaseHistoryRepository, itemRepisitory, gormDB)
	releaseController := controllers.NewReleaseController(releaseMessageService)
	dateController := controllers.NewDateController()
	initControllers := controllers.InitControllersFn(appNamespaceController, itemController, releaseHistoryController, releaseController, dateController)
	engine, err := http.NewRouter(httpOptions, logger, initControllers)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	application, err := admin.NewApp(adminOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, db.ProviderSet, controllers.ProviderSet, services.ProviderSet, repositories.ProviderSet, admin.ProviderSet, http.ProviderSet, httpclient.ProviderSet)
