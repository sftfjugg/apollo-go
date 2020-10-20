package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAddress,
	NewAppService,
	NewFavoriteService,
	NewAppNamespaceService,
	NewItemService,
	NewReleaseService,
)
