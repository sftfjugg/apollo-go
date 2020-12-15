package controllers

import (
	"encoding/json"
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
	writeNames := make([]string, 0)
	for _, w := range role.Write {
		writeNames = append(writeNames, w.UserName)
	}
	writerJson, err := json.Marshal(writeNames)
	if err != nil {
		c.String(http.StatusBadRequest, "bind write json error:%v", err)
		return
	}
	releaseNames := make([]string, 0)
	for _, w := range role.Release {
		releaseNames = append(releaseNames, w.UserName)
	}
	releaseJson, err := json.Marshal(releaseNames)
	if err != nil {
		c.String(http.StatusBadRequest, "bind write json error:%v", err)
		return
	}
	writes := string(writerJson)
	releases := string(releaseJson)
	query := "&writes=" + writes + "&releases=" + releases
	// 操作记录
	if role.Env == "TEST" {
		if c.Request.URL.RawQuery != "" {
			c.Request.URL.RawQuery += "&env_id=1&env_name=TEST" + query
		} else {
			c.Request.URL.RawQuery += "env_id=1&env_name=TEST" + query
		}
	} else if role.Env == "ALIYUN" {
		if c.Request.URL.RawQuery != "" {
			c.Request.URL.RawQuery += "&env_id=4&env_name=ALIYUN" + query
		} else {
			c.Request.URL.RawQuery += "env_id=4&env_name=ALIYUN" + query
		}
	} else {
		if c.Request.URL.RawQuery != "" {
			c.Request.URL.RawQuery += "&env_id=3&env_name=ONLINE" + query
		} else {
			c.Request.URL.RawQuery += "env_id=3&env_name=ONLINE" + query
		}
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

func (ctl RoleController) DeleteByUserId(c *gin.Context) {

	role := new(struct {
		UserId string `json:"user_id"`
	})
	if err := c.ShouldBind(role); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.DeleteByUserId(role.UserId); err != nil {
		c.String(http.StatusBadRequest, "call RoleService.CreateBackDoor failed", err)
		return
	}

}
