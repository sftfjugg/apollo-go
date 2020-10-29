package controllers

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAppController,
	NewAppNamespaceController,
	InitControllersFn,
	NewItemController,
	NewItemRelatedControllerr,
	NewAppNamespaceRelatedController,
	NewReleaseController)
