package controllers

import (
	"apollo-adminserivce/internal/pkg/http"
	"github.com/gin-gonic/gin"
)

func InitControllersFn(
	appNamespaceController *AppNamespaceController,
	itemController *ItemController,
	releaseController *ReleaseController,
) http.InitControllers {
	return func(r *gin.Engine) {

		{
			r.POST("/app_namespace", appNamespaceController.Create)
			r.PUT("/app_namespace", appNamespaceController.Update)
			r.GET("/app_namespace", appNamespaceController.FindAppNamespaceByAppIdAndClusterName)
			r.GET("/app_namespace/name", appNamespaceController.FindOneAppNamespaceByAppIdAndClusterNameAndName)
			r.GET("/app_namespace_all", appNamespaceController.FindAppNamespaceByAppId)
			r.DELETE("/app_namespace", appNamespaceController.DeleteById)
			r.DELETE("/app_namespace_by_name", appNamespaceController.DeleteByNameAndAppId)
			r.POST("/app_namespace_related", appNamespaceController.CreateByRelated)
		}
		{
			r.POST("/item", itemController.Create)
			r.POST("/items", itemController.Creates)
			r.PUT("/item", itemController.Update)
			r.DELETE("/items", itemController.DeleteByNamespaceId)
			r.GET("/items", itemController.FindItemByNamespaceId)
			r.GET("/items_by_key", itemController.FindItemByKeyForPage)
			r.GET("/items_by_key_on_app", itemController.FindAppItemByKeyForPage)
			r.DELETE("/item", itemController.DeleteById)
			r.GET("/item", itemController.FindItemByNamespaceIdAndKey)
			r.GET("/item_by_key_and_appId", itemController.FindItemByAppIdAndKey)
			r.GET("/items_release", itemController.FindItemByNamespaceIdOnRelease)
		}
		{
			r.POST("/release", releaseController.Create)
		}

	}

}
