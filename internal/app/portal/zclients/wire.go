package zclients

import "github.com/google/wire"

// ProviderSet wire provider of zclients zclients
var ProviderSet = wire.NewSet(NewUicService, NewLimosService, NewHttpClient, NewDepartmentService, NewUserService)
