package controllers

import (
	"apollo-adminserivce/internal/app/portal/services"
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
	env := c.Param("env")
	r, err := ctl.service.Create(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) CreateByRelated(c *gin.Context) {
	env := c.Param("env")
	param := new(struct {
		NamespaceId string `json:"namespace_id"`
		ClusterName string `json:"cluster_name"`
		AppId       string `json:"app_id"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind parm error:%v", err)
		return
	}
	r, err := ctl.service.CreateByRelated(param.NamespaceId, param.ClusterName, param.AppId, env)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.CreateByRelated run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) DeleteById(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.DeleteById(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.DeleteById run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
func (ctl AppNamespaceController) Update(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.Update(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Update run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) FindAppNamespaceByAppIdAndClusterName(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindAppNamespaceByAppIdAndClusterName(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppNamespaceByAppIdAndClusterName run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
