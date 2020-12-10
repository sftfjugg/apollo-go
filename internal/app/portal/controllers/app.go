package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"net/http"
)

type AppController struct {
	service services.AppService
}

func NewAppController(appService services.AppService) *AppController {
	return &AppController{service: appService}
}

//查看当前用户权限
func (ctl AppController) FindAuth(c *gin.Context) {
	userId, _ := c.Get("UserID")
	appId := c.GetHeader("AppId")
	cluster := c.Query("cluster")
	env := c.Query("env")
	auth, err := ctl.service.FindAuth(appId, userId.(string), cluster, env)
	if err != nil {
		c.String(http.StatusForbidden, "call app.FindAuth failed:%v", err)
		return
	}
	c.JSON(http.StatusOK, auth)
}

//查看所有分组
func (ctl AppController) FindGroupsOfDevelopment(c *gin.Context) {

	groups, err := ctl.service.FindGroupsOfDevelopment()
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindGroupsOfDevelopment() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

//查看limos所有项目
func (ctl AppController) FindLimosAppForPage(c *gin.Context) {
	param := new(struct {
		Name     string `form:"name"`
		Owner    string `form:"owner"`
		PageNum  int32  `form:"page_num"`
		PageSize int32  `form:"page_size"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindLimosAppForPage(param.Name, param.Owner, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindLimosAppForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}

//获得所有用户
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
