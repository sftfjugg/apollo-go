package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.didapinche.com/foundation/ophis"
	_ "go.didapinche.com/foundation/ophis"
	_ "go.didapinche.com/goapi/plat_operate_history_api"
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
	dingdingController *DingdingController,
	ophis *ophis.Api,
) http.InitControllers {
	return func(r *gin.Engine) {

		{
			r.GET("/cluster", uic.AuthLogin(), appNamespaceController.FindAllClusterNameByAppId)                                                   //查找集群
			r.POST("/cluster/:env", uic.AuthLogin(), ophis.OpenWriter(), appNamespaceController.CreateCluster, ophis.Record("apollo-plus-portal")) //创建集群
			r.GET("/app_namespace/:env", uic.AuthLogin(), appNamespaceController.FindAppNamespaceByAppIdAndClusterName)
			r.POST("/app_namespace/:env", uic.AuthLogin(), ophis.OpenWriter(), appNamespaceController.Create, ophis.Record("apollo-plus-portal"))                         //创建namespace
			r.POST("/app_namespace_by_lane/:env", uic.AuthLogin(), ophis.OpenWriter(), appNamespaceController.CreateLane, ophis.Record("apollo-plus-portal"))             //关联泳道
			r.DELETE("/app_namespace/:env", uic.AuthLogin(), ophis.OpenWriter(), appNamespaceController.DeleteById, ophis.Record("apollo-plus-portal"))                   //删除泳道
			r.DELETE("/app_namespace_by_name/:env", uic.AuthLogin(), ophis.OpenWriter(), appNamespaceController.DeleteByNameAndAppId, ophis.Record("apollo-plus-portal")) //删除集群
			r.PUT("/app_namespace/:env", uic.AuthLogin(), ophis.OpenWriter(), appNamespaceController.Update, ophis.Record("apollo-plus-portal"))                          //修改集群
			//r.PUT("/app_namespace_is_dispaly/:env", uic.AuthLogin(), appNamespaceController.UpdateIsDisply)
			r.GET("/app_namespace_all/:env", uic.AuthLogin(), appNamespaceController.FindAppNamespaceByAppId)
			r.GET("/app_by_lane", uic.AuthLogin(), appNamespaceController.FindByLaneName)
			r.GET("/app_by_app_and_lane", uic.AuthLogin(), appNamespaceController.FindAppByLaneNameandAppId)
		}
		{
			r.GET("/items/:env", uic.AuthLogin(), itemController.FindItemByNamespaceId)
			r.GET("/items_by_key/:env", uic.AuthLogin(), itemController.FindItemByKeyForPage)
			r.GET("/items_by_key_on_app/:env", uic.AuthLogin(), itemController.FindAppItemByKeyForPage)
			r.POST("/item/:env", uic.AuthLogin(), itemController.Create)
			r.POST("/item_by_text/:env", uic.AuthLogin(), itemController.CreateByText)
			r.DELETE("/item/:env", uic.AuthLogin(), itemController.DeleteById)
			r.PUT("/item/:env", uic.AuthLogin(), itemController.Update)
			r.PUT("/items/:env", uic.AuthLogin(), itemController.Updates) //批量更新
			r.GET("/item_comment/:env", uic.AuthLogin(), itemController.FindAllComment)
			r.DELETE("/items/:env", uic.AuthLogin(), itemController.DeleteByNamespaceId)
			r.GET("/item/:env", uic.AuthLogin(), itemController.FindItemByNamespaceIdAndKey)
			r.GET("/item_by_key_and_app_id/:env", uic.AuthLogin(), itemController.FindItemByAppIdAndKey)
		}
		{ //发布和灰度全量发布
			r.POST("/releases/:env", uic.AuthLogin(), releaseController.Creates) //批量发布
			r.POST("/release/:env", uic.AuthLogin(), ophis.OpenWriter(), releaseController.Create, ophis.Record("apollo-plus-portal"))
			r.POST("/release_gray_total/:env", uic.AuthLogin(), releaseController.ReleaseGrayTotal)
		}
		{ //历史
			r.GET("/release_history/:env", uic.AuthLogin(), historyController.Find)
		}
		{
			r.POST("/roles", uic.AuthLogin(), ophis.OpenWriter(), roleController.Create, ophis.Record("apollo-plus-portal")) //授权
			r.POST("/role_back_door", uic.AuthLogin(), roleController.CreateBackDoor)
			r.DELETE("/role_back_door", uic.AuthLogin(), roleController.DeleteByUserId)
			r.GET("/roles", uic.AuthLogin(), roleController.FindByAppId)
			r.GET("/auth", uic.AuthLogin(), controller.FindAuth)
			r.GET("/auths", uic.AuthLogin(), controller.FindAuths)
		}

		{
			r.GET("/dingding", uic.AuthLogin(), dingdingController.FindAll)
			r.POST("/dingding", uic.AuthLogin(), dingdingController.Create)
			r.PUT("/dingding", uic.AuthLogin(), dingdingController.Update)
			r.DELETE("/dingding", uic.AuthLogin(), dingdingController.Delete)

		}

		//limos相关
		{
			r.GET("/limos/apps_favorite", uic.AuthLogin(), controller.GetFavorite)
			r.GET("/limos/apps_owner", uic.AuthLogin(), controller.GetOwner)
			r.GET("/limos/apps_recent", uic.AuthLogin(), controller.GetRecent)
			r.GET("/limos/apps", uic.AuthLogin(), controller.FindLimosAppForPage)
			r.GET("/limos/groups", uic.AuthLogin(), controller.FindGroupsOfDevelopment)
			r.GET("/users", controller.GetAllUsers)
			//r.GET("/app/:appId", controller.FindByAppId)
		}

	}

}
