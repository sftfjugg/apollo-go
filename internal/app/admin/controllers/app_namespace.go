package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"net/http"
)

type AppNamespaceController struct {
	service services.AppNamespaceService
}

func NewAppNamespaceController(service services.AppNamespaceService) *AppNamespaceController {
	return &AppNamespaceController{service: service}
}

//创造集群
func (ctl AppNamespaceController) CreateCluster(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.ShouldBind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.CreateCluster error:%v")
		return
	}
	appNamespace.DataChange_CreatedBy = userId
	appNamespace.DataChange_LastModifiedBy = userId
	appNamespace.Format = "服务"
	appNamespace.LaneName = "default"
	appNamespace.Name = "application"
	appNamespace.IsPublic = false
	if err := ctl.service.Create(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespace)
}

func (ctl AppNamespaceController) Create(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.ShouldBind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v")
		return
	}
	appNamespace.DataChange_CreatedBy = userId
	appNamespace.DataChange_CreatedBy = userId
	if err := ctl.service.Create(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespace)
}

func (ctl AppNamespaceController) Update(c *gin.Context) {
	appNamespace := new(models.AppNamespace)
	if err := c.Bind(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v")
		return
	}
	appNamespace.DataChange_LastModifiedBy = userId
	if err := ctl.service.Update(appNamespace); err != nil {
		c.String(http.StatusBadRequest, "service.Update error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespace)
}

func (ctl AppNamespaceController) DeleteById(c *gin.Context) {
	appId := c.Query("id")
	if err := ctl.service.DeleteById(appId); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.DeleteById error:%v", err)
		return
	}
}

func (ctl AppNamespaceController) DeleteByNameAndAppIdAndCluster(c *gin.Context) {
	param := new(struct {
		AppId       string `form:"app_id" json:"app_id"`
		Name        string `form:"name" json:"name"`
		ClusterName string `form:"cluster_name" json:"cluster_name"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.DeleteByNameAndAppIdAndCluster(param.Name, param.AppId, param.ClusterName); err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.DeleteByNameAndAppId error:%v", err)
		return
	}
}

func (ctl AppNamespaceController) FindAppNamespaceByAppIdAndClusterName(c *gin.Context) {
	param := new(struct {
		AppId       string `form:"app_id" json:"app_id"`
		ClusterName string `form:"cluster_name" json:"cluster_name"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	appNamespaces, err := ctl.service.FindAppNamespaceByAppIdAndClusterName(param.AppId, param.ClusterName)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppNamespaceByAppIdAndClusterName error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespaces)
}

func (ctl AppNamespaceController) FindAppNamespace(c *gin.Context) {
	param := new(struct {
		AppId        string `form:"app_id"`
		Cluster_name string `form:"cluster_name"`
		Format       string `form:"format"`
		Comment      string `form:"comment"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	appNamespaces, err := ctl.service.FindAppNamespace(param.AppId, param.Cluster_name, param.Format, param.Comment)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAppNamespaceByAppId error:%v", err)
		return
	}
	c.JSON(http.StatusOK, appNamespaces)
}

func (ctl AppNamespaceController) FindAllClusterNameByAppId(c *gin.Context) {
	param := new(struct {
		AppId string `form:"app_id" json:"app_id"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	clusters, err := ctl.service.FindAllClusterNameByAppId(param.AppId)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindAllClusterNameByAppId error:%v", err)
		return
	}
	c.JSON(http.StatusOK, clusters)
}

func (ctl AppNamespaceController) FindByLaneName(c *gin.Context) {
	param := new(struct {
		Lane string `form:"lane" json:"lane"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	apps, err := ctl.service.FindByLaneName(param.Lane)
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.FindByLaneName error:%v", err)
		return
	}
	c.JSON(http.StatusOK, apps)
}
