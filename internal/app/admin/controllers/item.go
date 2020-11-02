package controllers

import (
	"github.com/gin-gonic/gin"
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

func (ctl ItemController) CreateOrUpdateItem(c *gin.Context) {
	item := new(models.Item)
	if err := c.ShouldBind(item); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.CreateOrUpdateItem(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.CreateOrUpdateItem() error:%v", err)
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
		c.String(http.StatusBadRequest, "AppNamespaceService.Create error:%v")
		return
	}
	item.DataChange_LastModifiedBy = userId
	if err := ctl.service.Update(item); err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.Create() error:%v", err)
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

	namespaceId := c.Query("app_id")
	comments, err := ctl.service.FindAllComment(namespaceId)
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
		AppId   string `form:"app_id"`
		Key     string `form:"key"`
		Format  string `form:"format"`
		Comment string `form:"comment"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindItemByAppIdAndKey(param.AppId, param.Key, param.Format, param.Comment)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByAppIdAndKey() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindItemByKeyForPage(c *gin.Context) {

	param := new(struct {
		Key      string `form:"key" json:"key"`
		Format   string `form:"format"`
		PageSize int    `form:"page_size"`
		PageNum  int    `form:"page_num"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindItemByKeyForPage(param.Key, param.Format, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindItemByKeyForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctl ItemController) FindAppItemByKeyForPage(c *gin.Context) {

	param := new(struct {
		Key      string `form:"key"`
		Format   string `form:"format"`
		PageSize int    `form:"page_size"`
		PageNum  int    `form:"page_num"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	items, err := ctl.service.FindAppItemByKeyForPage(param.Key, param.Format, param.PageSize, param.PageNum)
	if err != nil {
		c.String(http.StatusInternalServerError, "call ItemService.FindAppItemByKeyForPage() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, items)
}
