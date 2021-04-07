package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"net/http"
	"strconv"
)

type DingdingController struct {
	service services.DingdingService
}

func NewDingdingController(service services.DingdingService) *DingdingController {
	return &DingdingController{service: service}
}

func (ctl DingdingController) Create(c *gin.Context) {
	param := new(models.Dingding)
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Create(param); err != nil {
		c.String(http.StatusInternalServerError, "call DingdingController.Create error:%v", err)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func (ctl DingdingController) Update(c *gin.Context) {
	param := new(models.Dingding)
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Update(param); err != nil {
		c.String(http.StatusInternalServerError, "call DingdingController.Create error:%v", err)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func (ctl DingdingController) FindAll(c *gin.Context) {
	param := new(struct {
		PageSize int `form:"page_size"`
		PageNum  int `form:"page_num"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	dingdings, err := ctl.service.FindAll(param.PageNum, param.PageSize)
	if err != nil {
		c.String(http.StatusInternalServerError, "call DingdingController.Create error:%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  dingdings,
		"total": len(dingdings),
	})
}

func (ctl DingdingController) Delete(c *gin.Context) {
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Delete(ID); err != nil {
		c.String(http.StatusInternalServerError, "call DingdingController.Delete error:%v", err)
		return
	}
	c.JSON(http.StatusOK, 1)
}
