package controllers

import (
	"apollo-adminserivce/internal/app/adminservice/services"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemController struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (ctl ItemController) Create(c *gin.Context) {
	item := new(models.Item)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Create(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.Create() error:%v", err)
		return
	}
}

func (ctl ItemController) Update(c *gin.Context) {
	item := new(models.Item)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Update(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.Create() error:%v", err)
		return
	}
}

func (ctl ItemController) DeleteByNamespaceIdAndKey(c *gin.Context) {
	param := new(struct {
		NamespaceId string `form:"namespace_id",json:"namespace_id"`
		Key         string `form:"key",json:"key"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.DeleteByNamespaceIdAndKey(param.NamespaceId, param.Key); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.DeleteByNamespaceIdAndKey() error:%v", err)
		return
	}
}

func (ctl ItemController) DeleteByNamespaceId(c *gin.Context) {
	namespaceId := c.Query("namespace_id")
	if err := ctl.service.DeleteByNamespaceId(namespaceId); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.DeleteByNamespaceIdAndKey() error:%v", err)
		return
	}
}

func (ctl ItemController) FindItemByNamespaceId(c *gin.Context) {

	namespaceId := c.Query("namespace_id")
	items, err := ctl.service.FindItemByNamespaceId(namespaceId)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByNamespaceId() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindItemByNamespaceIdAndKey(c *gin.Context) {

	param := new(struct {
		NamespaceId string `form:"Namespace_id" json:"namespace_id"`
		Key         string `form:"Key" json:"key"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindItemByNamespaceIdAndKey(param.NamespaceId, param.Key)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByNamespaceIdAndKey() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}
