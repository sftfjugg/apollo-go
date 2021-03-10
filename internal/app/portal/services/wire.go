package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAddress,
	NewAppService,
	NewAppNamespaceService,
	NewItemService,
	NewReleaseService,
	NewZserviceApi,
	NewReleaseHistoryService,
	NewRoleService,
)
