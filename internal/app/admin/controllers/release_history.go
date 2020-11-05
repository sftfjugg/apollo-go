package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/services"
	"net/http"
)

type ReleaseHistoryController struct {
	service services.ReleaseHistoryService
}

func NewReleaseHistoryController(service services.ReleaseHistoryService) *ReleaseHistoryController {
	return &ReleaseHistoryController{service: service}
}

func (ctl ReleaseHistoryController) Find(c *gin.Context) {
	param := new(struct {
		Name     string `form:"name"`
		AppId    string `form:"app_id"`
		Key      string `form:"key"`
		PageSize int    `form:"page_size"`
		PageNum  int    `form:"page_num"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	releaseHistorys, err := ctl.service.Find(param.AppId, param.Name, param.Key, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusBadRequest, "call ReleaseHistoryService.Find:%v", err)
		return
	}
	c.JSON(http.StatusOK, releaseHistorys)
}
