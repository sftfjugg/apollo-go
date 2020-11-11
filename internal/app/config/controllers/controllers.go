package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
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
		{
			r.GET("/health", configController.Ping)
			r.GET("/ping", configController.Ping)
			r.GET("/Health", configController.Ping)
			r.GET("/Ping", configController.Ping)
			r.GET("/", configController.Ping)
		}
	}

}
