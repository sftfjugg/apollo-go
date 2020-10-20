package controllers

import (
	"apollo-adminserivce/internal/app/portal/services"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
}

func (ctl AppController) FindAllForPage(c *gin.Context) {
	param := new(struct {
		pageNum  int `form:"page_num"`
		PageSize int `form:"page_size"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindAllForPage(param.PageSize, param.pageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindAllForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}

func (ctl AppController) FindByNameOrAppIdForPage(c *gin.Context) {
	param := new(struct {
		name     string `from:"name"`
		appId    string `app_id`
		pageNum  int    `form:"page_num"`
		PageSize int    `form:"page_size"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindByNameOrAppIdForPage(param.name, param.appId, param.PageSize, param.pageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindByNameOrAppIdForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}

func (ctl AppController) FindByNameForPage(c *gin.Context) {
	param := new(struct {
		name     string `from:"name"`
		pageNum  int    `form:"page_num"`
		PageSize int    `form:"page_size"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindByNameForPage(param.name, param.PageSize, param.pageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindByNameForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}

func (ctl AppController) FindByAppId(c *gin.Context) {
	param := new(struct {
		appId string `uic:"app_id" binding:"required"`
	})
	if err := c.ShouldBindUri(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindByAppId(param.appId)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindByAppId() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}

func (ctl AppController) DeleteByAppId(c *gin.Context) {
	param := new(struct {
		appId string `app_id`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	err := ctl.service.DeleteByAppId(param.appId)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.DeleteByAppId() error:%v", err)
		return
	}
	c.String(http.StatusOK, "delete app successed")
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
		name     string `from:"name"`
		pageNum  int32  `form:"page_num"`
		PageSize int32  `form:"page_size"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindLimosAppForPage(param.name, param.PageSize, param.pageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindLimosAppForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}

func (ctl AppController) GetAllUsers(c *gin.Context) {
	param := new(struct {
		name string `from:"name"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	names, err := ctl.service.GetAllUsers(param.name)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.GetAllUsers() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, names)
}

func (ctl AppController) FindLimosAppById(c *gin.Context) {
	param := new(struct {
		appId int64 `app_id`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	app, err := ctl.service.FindLimosAppById(param.appId)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindLimosAppById() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, app)
}

func (ctl AppController) FindAuth(c *gin.Context) {
	param := new(struct {
		appId int64  `app_id`
		name  string `name`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	auth, err := ctl.service.FindAuth(param.appId, param.name)
	if err != nil {
		c.String(http.StatusInternalServerError, "call AppService.FindLimosAppById() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, auth)
}
