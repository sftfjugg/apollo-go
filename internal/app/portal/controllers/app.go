package controllers

import (
	"apollo-adminserivce/internal/app/portal/models"
	"apollo-adminserivce/internal/app/portal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AppController struct {
	service services.AppService
}

func NewAppController(appService services.AppService) *AppController {
	return &AppController{service: appService}
}

func (ctl AppController) Create(c *gin.Context) {

	app := new(models.App)
	if err := c.Bind(app); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Create(app); err != nil {
		c.String(http.StatusInternalServerError, "call AppService.Create() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, app)
}

func (ctl AppController) Update(c *gin.Context) {

	app := new(models.App)
	if err := c.Bind(app); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Update(app); err != nil {
		c.String(http.StatusInternalServerError, "call AppService.Update() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, app)
}

func (ctl AppController) DeleteByAppId(c *gin.Context) {

	appId := c.Query("app_id")
	if err := ctl.service.DeleteByAppId(appId); err != nil {
		c.String(http.StatusInternalServerError, "call AppService.DeleteByAppId() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appId)
}

func (ctl AppController) FindByName(c *gin.Context) {

	name := c.Query("name")
	app, err := ctl.service.FindByName(name)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindByName() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, app)
}

func (ctl AppController) FindByAppId(c *gin.Context) {

	appId := c.Query("app_id")
	app, err := ctl.service.FindByAppId(appId)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindByName() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, app)
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
