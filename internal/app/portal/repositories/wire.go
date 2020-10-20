package repositories

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewApplicationRepository, NewAppNamespaceRepository, NewFavoriteReposity)
