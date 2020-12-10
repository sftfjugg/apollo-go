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
	releaseController *ReleaseController,
	roleController *RoleController,
	historyController *ReleaseHistoryController,
) http.InitControllers {
	return func(r *gin.Engine) {

		{
			r.GET("/cluster", uic.AuthLogin(), appNamespaceController.FindAllClusterNameByAppId)
			r.POST("/cluster/:env", uic.AuthLogin(), appNamespaceController.CreateCluster)
			r.GET("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.FindAppNamespaceByAppIdAndClusterName)
			r.POST("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.Create)
			r.DELETE("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.DeleteById)
			r.DELETE("/app_namespace_by_name/:env", uic.AuthLogin(), appNamespaceController.DeleteByNameAndAppId)
			r.PUT("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.Update)
			//r.PUT("/app_namespace_is_dispaly/:env", uic.AuthLogin(), appNamespaceController.UpdateIsDisply)
			r.GET("/app_namespace_all/:env", uic.AuthLogin(), appNamespaceController.FindAppNamespaceByAppId)
			r.GET("/app_by_lane/:env", uic.AuthLogin(), appNamespaceController.FindByLaneName)
		}
		{
			r.GET("/items/:env", uic.AuthLogin(), itemController.FindItemByNamespaceId)
			r.GET("/items_by_key/:env", uic.AuthLogin(), itemController.FindItemByKeyForPage)
			r.GET("/items_by_key_on_app/:env", uic.AuthLogin(), itemController.FindAppItemByKeyForPage)
			r.POST("/item/:env", uic.AuthLogin(), itemController.Create)
			r.POST("/item_by_text/:env", uic.AuthLogin(), itemController.CreateByText)
			r.DELETE("/item/:env", uic.AuthLogin(), itemController.DeleteById)
			r.PUT("/item/:env", uic.AuthLogin(), itemController.Update)
			r.GET("/item_comment/:env", uic.AuthLogin(), itemController.FindAllComment)
			r.DELETE("/items/:env", uic.AuthLogin(), itemController.DeleteByNamespaceId)
			r.GET("/item/:env", uic.AuthLogin(), itemController.FindItemByNamespaceIdAndKey)
			r.GET("/item_by_key_and_app_id/:env", uic.AuthLogin(), itemController.FindItemByAppIdAndKey)
		}
		{ //发布和灰度全量发布
			r.POST("/release/:env", uic.AuthLogin(), releaseController.Create)
			r.POST("/release_gray_total/:env", uic.AuthLogin(), releaseController.ReleaseGrayTotal)
		}
		{ //历史
			r.GET("/release_history/:env", uic.AuthLogin(), historyController.Find)
		}
		{
			r.POST("/role", uic.AuthLogin(), roleController.Create)
			r.GET("/role", uic.AuthLogin(), roleController.FindByAppId)
			r.GET("/auth", uic.AuthLogin(), controller.FindAuth)
		}

		//权限相关，暂时保留2020.10.28 lihang
		{
			r.GET("/limos/apps", uic.AuthLogin(), controller.FindLimosAppForPage)
			r.GET("/limos/groups", uic.AuthLogin(), controller.FindGroupsOfDevelopment)
			r.GET("/users", controller.GetAllUsers)
			//r.GET("/app/:appId", controller.FindByAppId)
		}

	}

}
