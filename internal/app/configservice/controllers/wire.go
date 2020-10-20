package controllers

import "github.com/google/wire"

var ProviderSet = wire.NewSet(InitControllersFn, NewNotificationController, NewConsulController, NewConfigController)
