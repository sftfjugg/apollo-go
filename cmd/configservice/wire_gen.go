// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

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

// Injectors from wire.go:

func CreateApps(cf string) (*app.Application, error) {
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
	dbOptions, err := db.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	gormDB, err := db.New(dbOptions, logger)
	if err != nil {
		return nil, err
	}
	releaseMessageRepository := repositories.NewReleaseMessageRepository(gormDB)
	releaseMessageService := services.NewReleaseMessageService(releaseMessageRepository)
	configserviceOptions, err := configservice.NewOptions(viper, logger, releaseMessageService)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	configRepository := repositories.NewConfigRepository(gormDB)
	configService := services.NewConfigService(configRepository)
	configController := controllers.NewConfigController(configService)
	consulService := services.NewConsulService()
	consulController := controllers.NewConsulController(consulService)
	notificationMessageService := services.NewNotificationMessageService()
	notificationController := controllers.NewNotificationController(notificationMessageService)
	initControllers := controllers.InitControllersFn(configController, consulController, notificationController)
	engine, err := http.NewRouter(httpOptions, logger, initControllers)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	application, err := configservice.NewApp(configserviceOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSets = wire.NewSet(log.ProviderSet, config.ProviderSet, db.ProviderSet, repositories.ProviderSet, services.ProviderSet, controllers.ProviderSet, configservice.ProviderSet, http.ProviderSet, httpclient.ProviderSet)