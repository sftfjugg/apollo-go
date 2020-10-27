package controllers

import (
	"apollo-adminserivce/internal/pkg/http"
	"github.com/gin-gonic/gin"
	"go.didapinche.com/uic"
)

func InitControllersFn(
	controller *AppController,
	uic *uic.Api,
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
			r.DELETE("/app_namespace_by_name/:env", appNamespaceController.DeleteByNameAndAppId)
			r.PUT("/app_namespace/:env", appNamespaceController.Update)
			r.POST("/app_namespace_by_related/:env", appNamespaceController.CreateByRelated)
			r.GET("/app_namespace_all/:env", appNamespaceController.FindAppNamespaceByAppId)
		}
		{
			r.GET("/items/:env", itemController.FindItemByNamespaceId)
			r.GET("/items_by_key/:env", itemController.FindItemByKeyForPage)
			r.GET("/items_by_key_on_app/:env", itemController.FindAppItemByKeyForPage)
			r.POST("/item/:env", uic.AuthLogin(), itemController.Create)
			r.DELETE("/item/:env", uic.AuthLogin(), itemController.DeleteByNamespaceIdAndKey)
			r.PUT("/item/:env", uic.AuthLogin(), itemController.Update)
			r.DELETE("/items/:env", uic.AuthLogin(), itemController.DeleteByNamespaceId)
			r.GET("/item/:env", itemController.FindItemByNamespaceIdAndKey)
			r.GET("/item_by_key_and_app_id/:env", itemController.FindItemByAppIdAndKey)
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
			r.DELETE("/items_related", itemRelatedController.DeleteByNamespaceId)
			r.GET("/item_related", itemRelatedController.FindItemByNamespaceIdAndKey)
			r.GET("/items_related", itemRelatedController.FindItemByNamespaceId)
			r.GET("/item_related/key", itemRelatedController.FindOneItemByNamespaceIdAndKey)
		}

	}

}
