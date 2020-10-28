package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice/services"
	"net/http"
)

type ConsulController struct {
	services services.ConsulService
}

func NewConsulController(services services.ConsulService) *ConsulController {
	return &ConsulController{services: services}
}

func (ctl ConsulController) FindConfigService(c *gin.Context) {
	consul, err := ctl.services.FindAddress("config-service")
	if err != nil {
		c.String(http.StatusBadRequest, "call ConsulService.FindConsulByName error:%v", err)
		return
	}
	c.JSON(http.StatusOK, consul)
}

func (ctl ConsulController) FindAdminService(c *gin.Context) {
	consul, err := ctl.services.FindAddress("admin-service")
	if err != nil {
		c.String(http.StatusBadRequest, "call ConsulService.FindConsulByName error:%v", err)
		return
	}
	c.JSON(http.StatusOK, consul)
}
