package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.didapinche.com/uic"
)

func InitControllersFn(
	controller *AppController,
	uic *uic.Api,
	appNamespaceController *AppNamespaceController,
	itemController *ItemController,
	itemRelatedController *ItemRelatedController,
	releaseController *ReleaseController,
	historyController *ReleaseHistoryController,
	appNamespaceRelatedController *AppNamespaceRelatedController,
) http.InitControllers {
	return func(r *gin.Engine) {

		{
			r.GET("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.FindAppNamespaceByAppIdAndClusterName)
			r.POST("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.Create)
			r.DELETE("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.DeleteById)
			r.DELETE("/app_namespace_by_name/:env", uic.AuthLogin(), appNamespaceController.DeleteByNameAndAppId)
			r.PUT("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.Update)
			r.POST("/app_namespace_by_related/:env", uic.AuthLogin(), appNamespaceController.CreateByRelated)
			r.GET("/app_namespace_all/:env", uic.AuthLogin(), appNamespaceController.FindAppNamespaceByAppId)
		}
		{
			r.GET("/items/:env", uic.AuthLogin(), itemController.FindItemByNamespaceId)
			r.GET("/items_by_key/:env", uic.AuthLogin(), itemController.FindItemByKeyForPage)
			r.GET("/items_by_key_on_app/:env", uic.AuthLogin(), itemController.FindAppItemByKeyForPage)
			r.POST("/item/:env", uic.AuthLogin(), itemController.Create)
			r.DELETE("/item/:env", uic.AuthLogin(), itemController.DeleteById)
			r.PUT("/item/:env", uic.AuthLogin(), itemController.Update)
			r.GET("/item_comment/:env", uic.AuthLogin(), itemController.FindAllComment)
			r.DELETE("/items/:env", uic.AuthLogin(), itemController.DeleteByNamespaceId)
			r.GET("/item/:env", uic.AuthLogin(), itemController.FindItemByNamespaceIdAndKey)
			r.GET("/item_by_key_and_app_id/:env", uic.AuthLogin(), itemController.FindItemByAppIdAndKey)
		}
		{
			r.POST("/release/:env", uic.AuthLogin(), releaseController.Create)
			r.POST("/release_gray_total/:env", uic.AuthLogin(), releaseController.ReleaseGrayTotal)
		}
		{
			r.GET("/release_history/:env", uic.AuthLogin(), historyController.Find)
		}

		//权限相关，暂时保留2020.10.28 lihang
		{
			r.GET("/limos/apps", uic.AuthLogin(), controller.FindLimosAppForPage)
			r.GET("/limos/app", uic.AuthLogin(), controller.FindLimosAppById)
			r.GET("/limos/groups", uic.AuthLogin(), controller.FindGroupsOfDevelopment)
			r.GET("/limos/auth", uic.AuthLogin(), controller.FindAuth)
			r.GET("/users", controller.GetAllUsers)
			//r.GET("/app/:appId", controller.FindByAppId)
		}
		//关联获得配置，暂时不做2020.10.28 lihang
		//{
		//	r.POST("/app", controller.Create)
		//	r.PUT("/app", controller.Update)
		//	r.DELETE("/app", controller.DeleteByAppId)
		//	r.GET("/app", controller.FindByName)
		//	r.GET("/app_id", controller.FindByAppId)
		//	r.POST("/app_namespace_related", appNamespaceRelatedController.Create)
		//	r.GET("/app_namespace_related", appNamespaceRelatedController.FindAppNamespaceByNameForPage)
		//	r.DELETE("/app_namespace_related", appNamespaceRelatedController.Delete)
		//	r.PUT("/app_namespace_related", appNamespaceRelatedController.Update)
		//	r.GET("/app_namespace_related/name", appNamespaceRelatedController.FindAppNamespaceByName)
		//}
		//{
		//	r.POST("/item_related", itemRelatedController.Create)
		//	r.PUT("/item_related", itemRelatedController.Update)
		//	r.DELETE("/item_related", itemRelatedController.DeleteById)
		//	r.DELETE("/items_related", itemRelatedController.DeleteByNamespaceId)
		//	r.GET("/item_related", itemRelatedController.FindItemByNamespaceIdAndKey)
		//	r.GET("/items_related", itemRelatedController.FindItemByNamespaceId)
		//	r.GET("/item_related/key", itemRelatedController.FindOneItemByNamespaceIdAndKey)
		//}

	}

}
