package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/services"
	"net/http"
	"strconv"
)

type ReleaseController struct {
	service services.ReleaseMessageService
}

func NewReleaseController(service services.ReleaseMessageService) *ReleaseController {
	return &ReleaseController{service: service}
}

func (ctl ReleaseController) Create(c *gin.Context) {
	param := new(models.ReleaseRequest)
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if param.Operator == "" {
		userId, err := c.Cookie("UserID")
		if err != nil {
			c.String(http.StatusBadRequest, "UserID don't  null:%v")
			return
		}
		param.Operator = userId
	}
	if err := ctl.service.Create(param); err != nil {
		c.String(http.StatusInternalServerError, "call ReleaseMessageService.Create() error:%v", err)
		return
	}
}

//批量发布
func (ctl ReleaseController) Creates(c *gin.Context) {
	params := make([]*models.ReleaseRequest, 0)
	if err := c.Bind(&params); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "UserID don't  null:%v", err)
		return
	}
	for i, _ := range params {
		params[i].Operator = userId
	}
	if err := ctl.service.Creates(params); err != nil {
		c.String(http.StatusInternalServerError, "call ReleaseMessageService.Creates() error:%v", err)
		return
	}
}

//灰度全量发布
func (ctl ReleaseController) ReleaseGrayTotal(c *gin.Context) {
	param := new(struct {
		Name        string `json:"name"`
		AppId       string `json:"app_id"`
		NamespaceId int    `json:"namespace_id"`
		ClusterName string `json:"cluster_name"`
		Operator    string `json:"operator"`
		IsDeleted   bool   `json:"is_deleted"`
	})
	if err := c.Bind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if param.Operator == "" {
		userId, err := c.Cookie("UserID")
		if err != nil {
			c.String(http.StatusBadRequest, "UserID don't  null:%v")
			return
		}
		param.Operator = userId
	}
	if err := ctl.service.ReleaseGrayTotal(strconv.Itoa(param.NamespaceId), param.Name, param.AppId, param.ClusterName, param.Operator, param.IsDeleted); err != nil {
		c.String(http.StatusInternalServerError, "call ReleaseMessageService.ReleaseGrayTotal() error:%v", err)
		return
	}
}
