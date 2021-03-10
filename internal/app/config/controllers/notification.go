package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/services"
	"net/http"
)

type NotificationController struct {
	service services.NotificationMessageService
}

func NewNotificationController(service services.NotificationMessageService) *NotificationController {
	return &NotificationController{service: service}
}

func (ctl NotificationController) PollNotification(c *gin.Context) {
	param := new(struct {
		AppId         string `form:"appId"`
		Cluster       string `form:"cluster"`
		Notifications string `form:"notifications"`
		Ip            string `form:"ip"`
		DataCenter    string `form:"dataCenter"`
		Lane          string `form:"lane"`
	})
	if err := c.BindQuery(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if param.Lane == "" {
		param.Lane = "default"
	}
	//c.String(address.StatusNotModified,"")
	notifications, err := ctl.service.CompareV(param.AppId, param.Cluster, param.Notifications, param.Lane)
	if err != nil {
		c.String(http.StatusBadRequest, "CompareV failed:%v", err)
		return
	} else if notifications == nil {
		c.String(http.StatusNotModified, "No change in configuration file")
		return
	} else {
		c.JSON(http.StatusOK, notifications)
	}
}
