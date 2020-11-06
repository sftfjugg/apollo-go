package controllers

import (
	"github.com/gin-gonic/gin"
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
	param := new(struct {
		Name        string   `json:"name"`
		IsPublic    bool     `json:"is_public"`
		Comment     string   `json:"comment"`
		AppId       string   `json:"app_id"`
		ClusterName string   `json:"cluster_name"`
		NamespaceId uint64   `json:"namespace_id"`
		Keys        []string `json:"keys"`
		Operator    string   `json:"operator"`
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
	NamespaceId := strconv.Itoa(int(param.NamespaceId))
	if err := ctl.service.Create(param.AppId, param.ClusterName, param.Comment, param.Name, NamespaceId, param.Operator, param.IsPublic, param.Keys); err != nil {
		c.String(http.StatusInternalServerError, "call ReleaseMessageService.Create() error:%v", err)
		return
	}
}

//灰度全量发布
func (ctl ReleaseController) ReleaseGrayTotal(c *gin.Context) {
	param := new(struct {
		Name        string `json:"name"`
		AppId       string `json:"app_id"`
		NamespaceId uint64 `json:"namespace_id"`
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
	if err := ctl.service.ReleaseGrayTotal(string(param.NamespaceId), param.Name, param.AppId, param.Operator, param.IsDeleted); err != nil {
		c.String(http.StatusInternalServerError, "call ReleaseMessageService.ReleaseGrayTotal() error:%v", err)
		return
	}
}
