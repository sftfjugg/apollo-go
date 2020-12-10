package controllers

import (
	"github.com/gin-gonic/gin"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"net/http"
)

type RoleController struct {
	service services.RoleService
}

func NewRoleController(service services.RoleService) *RoleController {
	return &RoleController{service: service}
}

func (ctl RoleController) Create(c *gin.Context) {

	role := new(models2.Role)
	if err := c.ShouldBind(role); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Create(role); err != nil {
		c.String(http.StatusBadRequest, "call RoleService.Create failed", err)
		return
	}

}

//查看当前项目的编辑和发布用户
func (ctl RoleController) FindByAppId(c *gin.Context) {
	appId := c.GetHeader("AppId")
	cluster := c.Query("cluster")
	env := c.Query("env")
	name := c.Query("name")
	roles, err := ctl.service.FindByAppId(appId, cluster, env, name)
	if err != nil {
		c.String(http.StatusBadRequest, "call RoleService.Create failed", err)
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (ctl RoleController) CreateBackDoor(c *gin.Context) {

	role := new(struct {
		UserId string `json:"user_id"`
	})
	if err := c.ShouldBind(role); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.CreateBackDoor(role.UserId); err != nil {
		c.String(http.StatusBadRequest, "call RoleService.CreateBackDoor failed", err)
		return
	}

}
