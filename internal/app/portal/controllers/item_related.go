package controllers

import (
	"apollo-adminserivce/internal/app/portal/models"
	"apollo-adminserivce/internal/app/portal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemRelatedController struct {
	service services.ItemRelatedService
}

func NewItemRelatedControllerr(service services.ItemRelatedService) *ItemRelatedController {
	return &ItemRelatedController{service: service}
}

func (ctl ItemRelatedController) Create(c *gin.Context) {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	if err := ctl.service.Create(item); err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.Create error:%v", err)
		return
	}
}

func (ctl ItemRelatedController) Creates(c *gin.Context) {
	item := make([]*models.Item, 0)
	if err := c.Bind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	if err := ctl.service.Creates(item); err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.Creates error:%v", err)
		return
	}
}

func (ctl ItemRelatedController) Update(c *gin.Context) {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	if err := ctl.service.Update(item); err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.Update error:%v", err)
		return
	}
}

func (ctl ItemRelatedController) DeleteById(c *gin.Context) {
	id := c.Query("id")
	if err := ctl.service.DeleteById(id); err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.DeleteById error:%v", err)
		return
	}
}

func (ctl ItemRelatedController) DeleteByNamespaceId(c *gin.Context) {
	namespaceId := c.Query("namespace_id")
	if err := ctl.service.DeleteByNamespaceId(namespaceId); err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.DeleteByNamespaceId error:%v", err)
		return
	}
}

func (ctl ItemRelatedController) FindItemByNamespaceId(c *gin.Context) {
	namespaceId := c.Query("namespace_id")
	items, err := ctl.service.FindItemByNamespaceId(namespaceId)
	if err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.FindItemByNamespaceId error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)

}

func (ctl ItemRelatedController) FindItemByNamespaceIdAndKey(c *gin.Context) {
	param := new(struct {
		NamespaceId string `json:"namespace_id" form:"namespace_id"`
		Key         string `json:"key" form:"key"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	items, err := ctl.service.FindItemByNamespaceIdAndKey(param.NamespaceId, param.Key)
	if err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.FindItemByNamespaceIdAndKey error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)

}

func (ctl ItemRelatedController) FindOneItemByNamespaceIdAndKey(c *gin.Context) {
	param := new(struct {
		NamespaceId uint64 `json:"namespace_id" form:"namespace_id"`
		Key         string `json:"key" form:"key"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	item, err := ctl.service.FindOneItemByNamespaceIdAndKey(param.NamespaceId, param.Key)
	if err != nil {
		c.String(http.StatusBadRequest, "call ItemRelatedService.FindOneItemByNamespaceIdAndKey error:%v", err)
		return
	}
	c.JSON(http.StatusOK, item)

}
