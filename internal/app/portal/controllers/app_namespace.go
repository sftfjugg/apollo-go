package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
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
	UserID, _ := c.Get("UserID")
	if UserID.(string) == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.Create(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) CreateCluster(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	if UserID.(string) == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.CreateCluster(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.CreateCluster run failed:%v", err)
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

func (ctl AppNamespaceController) DeleteByNameAndAppId(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.DeleteByNameAndAppId(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.DeleteByNameAndAppId run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) Update(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	if UserID.(string) == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
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

func (ctl AppNamespaceController) FindAppNamespaceByAppId(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.FindAppNamespaceByAppId(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppNamespaceByAppId run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) FindAllClusterNameByAppId(c *gin.Context) {
	r, err := ctl.service.FindAllClusterNameByAppId(c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAllClusterNameByAppId run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
