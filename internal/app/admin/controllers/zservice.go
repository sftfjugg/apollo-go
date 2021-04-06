package controllers

import (
	"github.com/gin-gonic/gin"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"net/http"
)

type ZServiceController struct {
	service services.ZService
}

func NewZServiceController(service services.ZService) *ZServiceController {
	return &ZServiceController{service: service}
}

func (ctl ZServiceController) CreateOrFindAppNamespace(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.ShouldBind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	id, err := ctl.service.CreateOrFindAppNamespace(appNamespace)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.CreateOrFindAppNamespace error:%v", err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func (ctl ZServiceController) CreateOrUpdateItem(c *gin.Context) {
	item := new(models.Item)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.CreateOrUpdateItem(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.CreateOrUpdateItem() error:%v", err)
		return
	}
}

func (ctl ZServiceController) PublishNamespace(c *gin.Context) {
	param := new(models2.ReleaseRequest)
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.PublishNamespace(param); err != nil {
		c.String(http.StatusInternalServerError, "call ReleaseMessageService.Create() error:%v", err)
		return
	}
}
