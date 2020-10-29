package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewConfigService, NewReleaseMessageService, NewConsulService, NewNotificationMessageService)
