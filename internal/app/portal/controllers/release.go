package controllers

import (
	"apollo-adminserivce/internal/app/portal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReleaseController struct {
	service services.ReleaseService
}

func NewReleaseController(service services.ReleaseService) *ReleaseController {
	return &ReleaseController{service: service}
}

func (ctl ReleaseController) Create(c *gin.Context) {
	env := c.Param("env")
	r, err := ctl.service.Create(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ReleaseController.Create run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
