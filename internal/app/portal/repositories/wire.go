package repositories

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAppNamespaceRelatedRepository,
	NewAppRepository,
	NewItemRelatedRepisitory,
)
