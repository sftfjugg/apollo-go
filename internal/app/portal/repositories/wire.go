package repositories

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewRoleReposotory, NewHistoryRepository, NewDingdingRepository)
