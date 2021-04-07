package controllers

import (
	"github.com/gin-gonic/gin"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"net/http"
)

type ItemController struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (ctl ItemController) Create(c *gin.Context) {
	item := new(models.Item)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v")
		return
	}
	item.DataChange_CreatedBy = userId
	item.DataChange_LastModifiedBy = userId
	if err := ctl.service.Create(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.Create() error:%v", err)
		return
	}
}

func (ctl ItemController) CreateByText(c *gin.Context) {
	item := new(models2.ItemText)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "UserID is null error:%v")
		return
	}
	item.Operator = userId
	if err := ctl.service.CreateByText(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.CreateByText() error:%v", err)
		return
	}
}

func (ctl ItemController) Creates(c *gin.Context) {
	item := make([]*models.Item, 0)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Creates(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.Creates() error:%v", err)
		return
	}
}

func (ctl ItemController) Update(c *gin.Context) {
	item := new(models.Item)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "don't userId")
		return
	}
	item.DataChange_LastModifiedBy = userId
	if err := ctl.service.Update(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.Update() error:%v", err)
		return
	}
}

//批量更新
func (ctl ItemController) Updates(c *gin.Context) {
	items := make([]*models.Item, 0)
	if err := c.ShouldBind(items); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "don't userId")
		return
	}
	for i, _ := range items {
		items[i].DataChange_LastModifiedBy = userId
	}
	if err := ctl.service.Updates(items); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.Updates() error:%v", err)
		return
	}
}

func (ctl ItemController) DeleteById(c *gin.Context) {
	id := c.Query("id")
	userId, err := c.Cookie("UserID")
	if err != nil {
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v")
		return
	}
	if err := ctl.service.DeleteById(id, userId); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.DeleteById() error:%v", err)
		return
	}
}

func (ctl ItemController) DeleteByNamespaceId(c *gin.Context) {
	namespaceId := c.Query("namespace_id")
	if err := ctl.service.DeleteByNamespaceId(namespaceId); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.DeleteByNamespaceIdAndKey() error:%v", err)
		return
	}
}

func (ctl ItemController) FindItemByNamespaceId(c *gin.Context) {

	namespaceId := c.Query("namespace_id")
	items, err := ctl.service.FindItemByNamespaceId(namespaceId, "")
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByNamespaceId() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindAllComment(c *gin.Context) {

	appId := c.Query("app_id")
	name := c.Query("name")
	comments, err := ctl.service.FindAllComment(appId, name)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindAllComment() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (ctl ItemController) FindItemByNamespaceIdOnRelease(c *gin.Context) {

	namespaceId := c.Query("namespace_id")
	items, err := ctl.service.FindItemByNamespaceIdOnRelease(namespaceId)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByNamespaceIdOnRelease() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindItemByNamespaceIdAndKey(c *gin.Context) {

	param := new(struct {
		NamespaceId string `form:"namespace_id"`
		Key         string `form:"key"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindItemByNamespaceIdAndKey(param.NamespaceId, param.Key)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByNamespaceIdAndKey() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindItemByAppIdAndKey(c *gin.Context) {

	param := new(struct {
		AppId        string `form:"app_id"`
		Key          string `form:"key"`
		Cluster_name string `form:"cluster_name"`
		Format       string `form:"format"`
		Comment      string `form:"comment"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindItemByAppIdAndKey(param.AppId, param.Cluster_name, param.Key, param.Format, param.Comment)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByAppIdAndKey() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindItemByKeyForPage(c *gin.Context) {

	param := new(struct {
		Key          string `form:"key" json:"key"`
		Format       string `form:"format"`
		Cluster_name string `form:"cluster_name"`
		Comment      string `form:"comment"`
		PageSize     int    `form:"page_size"`
		PageNum      int    `form:"page_num"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindItemByKeyForPage(param.Cluster_name, param.Key, param.Format, param.Comment, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByKeyForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindAppItemByKeyForPage(c *gin.Context) {

	param := new(struct {
		Key          string `form:"key"`
		Format       string `form:"format"`
		Comment      string `form:"comment"`
		Cluster_name string `form:"cluster_name"`
		PageSize     int    `form:"page_size"`
		PageNum      int    `form:"page_num"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindAppItemByKeyForPage(param.Cluster_name, param.Key, param.Format, param.Comment, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindAppItemByKeyForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindItemByAppIdLikeKey(c *gin.Context) {
	appId := c.Query("app_id")
	key := c.Query("key")
	items, err := ctl.service.FindItemByAppIdLikeKey(appId, key)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByAppIdLikeKey() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}
