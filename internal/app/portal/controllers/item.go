package controllers

import (
	"apollo-adminserivce/internal/app/portal/services"
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
	env := c.Param("env")
	r, err := ctl.service.Create(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.Create run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) Update(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.Update(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.Update run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) DeleteByNamespaceIdAndKey(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.DeleteByNamespaceIdAndKey(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.DeleteByNamespaceIdAndKey run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) DeleteByNamespaceId(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.DeleteByNamespaceId(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.DeleteByNamespaceId run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) FindItemByNamespaceId(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindItemByNamespaceId(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.FindItemByNamespaceId run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) FindItemByNamespaceIdAndKey(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindItemByNamespaceIdAndKey(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.FindItemByNamespaceIdAndKey run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
