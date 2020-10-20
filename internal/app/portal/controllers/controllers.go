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
	releaseController *ReleaseController,
) http.InitControllers {
	return func(r *gin.Engine) {
		{
			r.POST("/app", controller.Create)
			r.PUT("/app", controller.Update)
			r.DELETE("/app", controller.DeleteByAppId)
			r.GET("/apps", controller.FindAllForPage)
			r.GET("/limos/apps", controller.FindLimosAppForPage)
			r.GET("/limos/app", controller.FindLimosAppById)
			r.GET("/limos/groups", controller.FindGroupsOfDevelopment)
			r.GET("/limos/auth", controller.FindAuth)
			r.GET("/apps/by", controller.FindByNameOrAppIdForPage)
			r.GET("/users", controller.GetAllUsers)
			//r.GET("/app/:appId", controller.FindByAppId)
			r.GET("apps/name", controller.FindByNameForPage)
		}

		{
			r.GET("/namespace/:env", appNamespaceController.FindAppNamespaceByAppIdAndClusterName)
			r.POST("/namespace/:env", appNamespaceController.Create)
			r.DELETE("/namespace/:env", appNamespaceController.DeleteById)
			r.PUT("/namespace/:env", appNamespaceController.Update)
		}
		{
			r.GET("/items/:env", itemController.FindItemByNamespaceId)
			r.POST("/item/:env", itemController.Create)
			r.DELETE("/item/:env", itemController.DeleteByNamespaceId)
			r.PUT("/item/:env", itemController.Update)
			r.DELETE("/items/:env", itemController.DeleteByNamespaceIdAndKey)
			r.GET("/item/:env", itemController.FindItemByNamespaceIdAndKey)
		}
		{
			r.POST("/release/:env", releaseController.Create)
		}

	}

}
