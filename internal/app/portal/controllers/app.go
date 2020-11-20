package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"net/http"
	"strconv"
)

type AppController struct {
	service services.AppService
}

func NewAppController(appService services.AppService) *AppController {
	return &AppController{service: appService}
}

func (ctl AppController) FindGroupsOfDevelopment(c *gin.Context) {

	groups, err := ctl.service.FindGroupsOfDevelopment()
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindGroupsOfDevelopment() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (ctl AppController) FindLimosAppForPage(c *gin.Context) {
	param := new(struct {
		Name     string `form:"name"`
		PageNum  int32  `form:"page_num"`
		PageSize int32  `form:"page_size"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindLimosAppForPage(param.Name, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindLimosAppForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}

func (ctl AppController) GetAllUsers(c *gin.Context) {
	param := new(struct {
		Name string `form:"name"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	names, err := ctl.service.GetAllUsers(param.Name)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.GetAllUsers() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, names)
}

func (ctl AppController) FindLimosAppById(c *gin.Context) {
	param := new(struct {
		AppId int64 `form:"app_id""`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	app, err := ctl.service.FindLimosAppById(param.AppId)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindLimosAppById() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, app)
}

func (ctl AppController) FindAuth(c *gin.Context) {

	id := c.GetHeader("AppId")
	Name := c.GetHeader("UserName")
	appId, err := strconv.ParseInt(id, 10, 64)
	auth, err := ctl.service.FindAuth(appId, Name)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindLimosAppById() error:%v", err)
		return
	}
	if !auth {
		c.AbortWithStatus(http.StatusForbidden)
		c.String(http.StatusForbidden, "call AppService.FindLimosAppById() error:%v")
		return
	}
}
