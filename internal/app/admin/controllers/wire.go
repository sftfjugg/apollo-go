package controllers

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAppNamespaceController,
	NewItemController,
	NewReleaseController,
	NewReleaseHistoryController,
	InitControllersFn,
	NewDateController,
	NewZServiceController,
)
