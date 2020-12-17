package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewReleaseMessageService,
	NewAppNamespaceService,
	NewItemService,
	NewReleaseHistoryService,
	NewZService,
)
