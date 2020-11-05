package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"net/http"
)

type ReleaseHistoryController struct {
	service services.ReleaseHistoryService
}

func NewReleaseHistoryController(service services.ReleaseHistoryService) *ReleaseHistoryController {
	return &ReleaseHistoryController{service: service}
}

func (ctl ReleaseHistoryController) Find(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.Find(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ReleaseHistoryController.Find run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
