package controllers

import (
	"apollo-adminserivce/internal/pkg/http"
	"github.com/gin-gonic/gin"
)

//r.GET("/services/config",consul.FindConfig)

func InitControllersFn(
	configController *ConfigController,
	consulController *ConsulController,
	notificationController *NotificationController,
) http.InitControllers {
	return func(r *gin.Engine) {
		{
			r.GET("/services/config", consulController.FindConfigService)
			r.GET("/services/admin", consulController.FindAdminService)

		}
		{
			r.GET("/configs/:appId/:clusterName/:namespace", configController.FindConfig)
		}
		{
			r.GET("/notifications/v2", notificationController.PollNotification)
		}
	}

}
