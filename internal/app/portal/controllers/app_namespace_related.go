package controllers

import (
	"apollo-adminserivce/internal/app/portal/models"
	"apollo-adminserivce/internal/app/portal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppNamespaceRelatedController struct {
	service services.AppNamespaceRelatedService
}

func NewAppNamespaceRelatedController(service services.AppNamespaceRelatedService) *AppNamespaceRelatedController {
	return &AppNamespaceRelatedController{service: service}
}

func (ctl AppNamespaceRelatedController) Create(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.Bind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	if err := ctl.service.Create(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceRelatedService.Create run failed:%v", err)
		return
	}
}

func (ctl AppNamespaceRelatedController) Delete(c *gin.Context) {
	id := c.Query("id")
	if err := ctl.service.Delete(id); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceRelatedService.Delete run failed:%v", err)
		return
	}
}

func (ctl AppNamespaceRelatedController) Update(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.Bind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	if err := ctl.service.Update(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceRelatedService.Update run failed:%v", err)
		return
	}
}

func (ctl AppNamespaceRelatedController) FindAppNamespaceByName(c *gin.Context) {
	name := c.Query("name")
	appNamespace, err := ctl.service.FindAppNamespaceByName(name)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceRelatedService.FindAppNamespaceByName run failed:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespace)
}

func (ctl AppNamespaceRelatedController) FindAppNamespaceByNameForPage(c *gin.Context) {
	param := new(struct {
		Name     string `json:"name" form:"name"`
		PageSize int    `json:"page_size" form:"page_size"`
		PageNum  int    `json:"page_num" form:"page_num"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
	}
	appNamespaces, err := ctl.service.FindAppNamespaceByNameForPage(param.Name, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceRelatedService.FindAppNamespaceByNameForPage run failed:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespaces)
}

func (ctl AppNamespaceRelatedController) FindAppNamespaceByDepartmentForPage(c *gin.Context) {
	param := new(struct {
		Department string `json:"department" form:"department"`
		PageSize   int    `json:"page_size" form:"page_size"`
		PageNum    int    `json:"page_num" form:"page_num"`
	})
	appNamespaces, err := ctl.service.FindAppNamespaceByDepartmentForPage(param.Department, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceRelatedService.FindAppNamespaceByDepartmentForPage run failed:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespaces)
}
