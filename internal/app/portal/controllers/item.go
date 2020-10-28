package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
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
	UserID, _ := c.Get("UserID")
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.Create(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.Create run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) Update(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.Update(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.Update run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) DeleteById(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.DeleteById(env, c.Request)
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

func (ctl ItemController) FindItemByNamespaceIdOnRelease(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindItemByNamespaceIdOnRelease(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.FindItemByNamespaceIdOnRelease run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) FindItemByKeyForPage(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindItemByKeyForPage(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.FindItemByKeyForPage run failed:%v", err)
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

func (ctl ItemController) FindItemByAppIdAndKey(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindItemByAppIdAndKey(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.FindItemByAppIdAndKey run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ItemController) FindAppItemByKeyForPage(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindAppItemByKeyForPage(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ItemController.FindAppItemByKeyForPage run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
