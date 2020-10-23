package controllers

import (
	"apollo-adminserivce/internal/app/adminservice/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReleaseController struct {
	service services.ReleaseMessageService
}

func NewReleaseController(service services.ReleaseMessageService) *ReleaseController {
	return &ReleaseController{service: service}
}

func (ctl ReleaseController) Create(c *gin.Context) {
	param := new(struct {
		Name          string   `json:"name"`
		IsPublic      bool     `json:"is_public"`
		Comment       string   `json:"comment"`
		AppId         string   `json:"app_id"`
		ClusterName   string   `json:"cluster_name"`
		NamespaceName string   `json:"namespace_name"`
		NamespaceId   string   `json:"namespace_id"`
		Keys          []string `json:"keys"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Create(param.Name, param.AppId, param.ClusterName, param.Comment, param.NamespaceName, param.NamespaceId, param.IsPublic, param.Keys); err != nil {
		c.String(http.StatusInternalServerError, "call ReleaseMessageService.Create() error:%v", err)
		return
	}
}
