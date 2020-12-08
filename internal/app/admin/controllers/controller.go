package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
)

func InitControllersFn(
	appNamespaceController *AppNamespaceController,
	itemController *ItemController,
	releaseHistoryController *ReleaseHistoryController,
	releaseController *ReleaseController,
	dataController *DateController, //导数据
) http.InitControllers {
	return func(r *gin.Engine) {

		{
			r.GET("/cluster", appNamespaceController.FindAllClusterNameByAppId)
			r.POST("/cluster", appNamespaceController.CreateCluster)
			r.POST("/app_namespace", appNamespaceController.Create)
			r.POST("/app_namespace/create_or_find", appNamespaceController.CreateOrFindAppNamespace)
			r.PUT("/app_namespace", appNamespaceController.Update)
			r.PUT("/app_namespace_is_dispaly", appNamespaceController.UpdateIsDisply)
			r.GET("/app_namespace", appNamespaceController.FindAppNamespaceByAppIdAndClusterName)
			r.GET("/app_namespace_all", appNamespaceController.FindAppNamespace)
			r.GET("/app_by_lane", appNamespaceController.FindByLaneName)
			r.DELETE("/app_namespace", appNamespaceController.DeleteById)
			r.DELETE("/app_namespace_by_name", appNamespaceController.DeleteByNameAndAppIdAndCluster)
		}
		{
			r.POST("/item", itemController.Create)
			r.POST("/item_by_text", itemController.CreateByText)
			r.POST("/item/create_or_update", itemController.CreateOrUpdateItem)
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
			r.GET("/item_comment", itemController.FindAllComment)
		}
		{
			r.POST("/release", releaseController.Create)
			r.POST("/release_gray_total", releaseController.ReleaseGrayTotal)
		}
		{
			r.GET("/release_history", releaseHistoryController.Find)
		}
		//{
		//	r.POST("/import_data", dataController.ImportDate) //导数据
		//	r.POST("/update_data", dataController.UpdateDate) //导数据
		//}
		{
			r.POST("/health", dataController.Health)
			r.GET("/health", dataController.Health)
		}
	}

}
