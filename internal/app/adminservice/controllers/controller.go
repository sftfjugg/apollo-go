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
			r.DELETE("/app_namespace", appNamespaceController.DeleteById)
			r.POST("/app_namespace_related,", appNamespaceController.CreateByRelated)
		}
		{
			r.POST("/item", itemController.Create)
			r.POST("/items", itemController.Creates)
			r.PUT("/item", itemController.Update)
			r.DELETE("/items", itemController.DeleteByNamespaceId)
			r.GET("/items", itemController.FindItemByNamespaceId)
			r.DELETE("/item", itemController.DeleteById)
			r.GET("/item", itemController.FindItemByNamespaceIdAndKey)
		}
		{
			r.POST("/release", releaseController.Create)
		}

	}

}
