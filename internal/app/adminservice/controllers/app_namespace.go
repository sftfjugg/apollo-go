package controllers

import (
	"apollo-adminserivce/internal/app/adminservice/services"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppNamespaceController struct {
	service services.AppNamespaceService
}

func NewAppNamespaceController(service services.AppNamespaceService) *AppNamespaceController {
	return &AppNamespaceController{service: service}
}

func (ctl AppNamespaceController) Create(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.ShouldBind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Create(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespace)
}

func (ctl AppNamespaceController) CreateByRelated(c *gin.Context) {
	param := new(struct {
		AppNamespace *models.AppNamespace `json:"app_namespace"`
		Items        []*models.Item       `json:"items"`
		AppName      string               `json:"app_name"`
		AppId        string               `json:"app_id"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.CreateByRelated(param.AppNamespace, param.Items, param.AppName, param.AppId); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v", err)
		return
	}
}

func (ctl AppNamespaceController) Update(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.Bind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Update(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Update error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespace)
}

func (ctl AppNamespaceController) DeleteById(c *gin.Context) {
	appId := c.Query("id")
	if err := ctl.service.DeleteById(appId); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.DeleteById error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appId)
}

func (ctl AppNamespaceController) FindAppNamespaceByAppIdAndClusterName(c *gin.Context) {
	param := new(struct {
		AppId       string `form:"app_id" json:"app_id"`
		ClusterName string `form:"cluster_name" json:"cluster_name"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	appNamespaces, err := ctl.service.FindAppNamespaceByAppIdAndClusterName(param.AppId, param.ClusterName)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppNamespaceByAppIdAndClusterName error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespaces)
}

func (ctl AppNamespaceController) FindAppNamespaceByAppId(c *gin.Context) {
	appId := c.Query("app_id")
	appNamespaces, err := ctl.service.FindAppNamespaceByAppId(appId)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppNamespaceByAppId error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespaces)
}

func (ctl AppNamespaceController) FindOneAppNamespaceByAppIdAndClusterNameAndName(c *gin.Context) {
	param := new(struct {
		Name        string `json:"name" form:"name"`
		AppId       string `form:"app_id" json:"app_id"`
		ClusterName string `form:"cluster_name" json:"cluster_name"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	appNamespace, err := ctl.service.FindOneAppNamespaceByAppIdAndClusterNameAndName(param.AppId, param.ClusterName, param.Name)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindOneAppNamespaceByAppIdAndClusterNameAndName error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespace)
}
