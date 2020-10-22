package controllers

import (
	"apollo-adminserivce/internal/pkg/http"
	"github.com/gin-gonic/gin"
)

func InitControllersFn(
	controller *AppController,
	//uic *uic.Api,
	appNamespaceController *AppNamespaceController,
	itemController *ItemController,
	itemRelatedController *ItemRelatedController,
	releaseController *ReleaseController,
	appNamespaceRelatedController *AppNamespaceRelatedController,
) http.InitControllers {
	return func(r *gin.Engine) {
		{
			r.POST("/app", controller.Create)
			r.PUT("/app", controller.Update)
			r.DELETE("/app", controller.DeleteByAppId)
			r.GET("/app", controller.FindByName)
			r.GET("/app_id", controller.FindByAppId)
			r.GET("/limos/apps", controller.FindLimosAppForPage)
			r.GET("/limos/app", controller.FindLimosAppById)
			r.GET("/limos/groups", controller.FindGroupsOfDevelopment)
			r.GET("/limos/auth", controller.FindAuth)
			r.GET("/users", controller.GetAllUsers)
			//r.GET("/app/:appId", controller.FindByAppId)
		}

		{
			r.GET("/app_namespace/:env", appNamespaceController.FindAppNamespaceByAppIdAndClusterName)
			r.POST("/app_namespace/:env", appNamespaceController.Create)
			r.DELETE("/app_namespace/:env", appNamespaceController.DeleteById)
			r.PUT("/app_namespace/:env", appNamespaceController.Update)
			r.POST("/app_namespace/related/:env", appNamespaceController.CreateByRelated)
		}
		{
			r.GET("/items/:env", itemController.FindItemByNamespaceId)
			r.POST("/item/:env", itemController.Create)
			r.DELETE("/item/:env", itemController.DeleteByNamespaceIdAndKey)
			r.PUT("/item/:env", itemController.Update)
			r.DELETE("/items/:env", itemController.DeleteByNamespaceId)
			r.GET("/item/:env", itemController.FindItemByNamespaceIdAndKey)
		}
		{
			r.POST("/release/:env", releaseController.Create)
		}
		{
			r.POST("/app_namespace_related", appNamespaceRelatedController.Create)
			r.GET("/app_namespace_related", appNamespaceRelatedController.FindAppNamespaceByNameForPage)
			r.DELETE("/app_namespace_related", appNamespaceRelatedController.Delete)
			r.PUT("/app_namespace_related", appNamespaceRelatedController.Update)
			r.GET("/app_namespace_related/name", appNamespaceRelatedController.FindAppNamespaceByName)
		}
		{
			r.POST("/item_related", itemRelatedController.Create)
			r.PUT("/item_related", itemRelatedController.Update)
			r.DELETE("/item_related", itemRelatedController.DeleteById)
			r.GET("/item_related", itemRelatedController.FindItemByNamespaceIdAndKey)
			r.GET("/items_related", itemRelatedController.FindItemByNamespaceId)
			r.GET("/item_related/key", itemRelatedController.FindOneItemByNamespaceIdAndKey)
		}

	}

}
