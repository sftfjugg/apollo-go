package repositories

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAppNamespaceRepository,
	NewItemRepisitory,
	NewRelease,
	NewReleaseMessageRepository,
	NewReleaseHistoryRepository,
)
