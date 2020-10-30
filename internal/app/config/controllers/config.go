package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/services"
	"net/http"
)

type ConfigController struct {
	service services.ConfigService
}

func NewConfigController(service services.ConfigService) *ConfigController {
	return &ConfigController{service: service}
}

func (ctl ConfigController) FindConfig(c *gin.Context) {
	param := new(struct {
		AppId       string `uri:"appId" binding:"required"`
		ClusterName string `uri:"clusterName" binding:"required"`
		Namespace   string `uri:"namespace" binding:"required"`
		DataCenter  string `form:"dataCenter"`
		ReleaseKey  string `form:"releaseKey"`
		Ip          string `form:"ip"`
		Messages    string `form:"messages"`
	})
	if err := c.ShouldBindUri(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	config, err := ctl.service.FindConfigByAppIdandCluster(param.AppId, param.ClusterName, param.Namespace)
	config.NamespaceName = param.Namespace
	if err != nil {
		c.String(http.StatusBadRequest, "get Config failed:%v", err)
		return
	}
	c.JSON(http.StatusOK, config)

}
