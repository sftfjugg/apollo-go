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
		c.String(http.StatusBadRequest, "AppNamespaceService.Create run failed", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) DeleteById(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.DeleteById(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.DeleteById run failed", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
func (ctl AppNamespaceController) Update(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.Update(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Update run failed", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) FindAppNamespaceByAppIdAndClusterName(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindAppNamespaceByAppIdAndClusterName(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppNamespaceByAppIdAndClusterName run failed", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
