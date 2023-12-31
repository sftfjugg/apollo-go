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

func (ctl AppNamespaceController) CreateLane(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	if UserID.(string) == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.CreateLane(env, c.Request)
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

func (ctl AppNamespaceController) UpdateIsDisply(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	if UserID.(string) == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.UpdateIsDisply(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.UpdateIdDisply run failed:%v", err)
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
	UserID, _ := c.Get("UserID")
	userId := UserID.(string)
	if userId == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	UserName, _ := c.Get("UserName")
	userName := UserName.(string)
	if userName == "" {
		c.String(http.StatusUnauthorized, "UserName don't null")
		return
	}
	appId := c.Query("app_id")
	r, err := ctl.service.FindAppNamespaceByAppId(env, userId, userName, appId, c.Request)
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

func (ctl AppNamespaceController) FindByLaneName(c *gin.Context) {
	r, err := ctl.service.FindByLaneName(c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindByLaneName run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl AppNamespaceController) FindAppByLaneNameandAppId(c *gin.Context) {
	r, err := ctl.service.FindAppByLaneNameandAppId(c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppByLaneNameandAppId run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
