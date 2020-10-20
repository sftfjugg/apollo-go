package controllers

import (
	"apollo-adminserivce/internal/app/portal/services"
	"apollo-adminserivce/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteController struct {
	service services.FavoriteService
}

func NewFavoriteController(service services.FavoriteService) *FavoriteController {
	return &FavoriteController{service: service}
}

func (ctl FavoriteController) Create(c *gin.Context) {
	favorite := new(models.Favorite)
	if err := c.BindJSON(favorite); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.Create(favorite); err != nil {
		c.String(http.StatusInternalServerError, "call FavoriteRepository.Create() error:%v", err)
		return
	}
}

func (ctl FavoriteController) DeleteByUserIdAndAppId(c *gin.Context) {
	params := new(struct {
		AppId  string `form:"app_id"`
		UserId string `form:"user_id"`
	})
	if err := c.BindQuery(params); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	if err := ctl.service.DeleteByUserIdAndAppId(params.UserId, params.AppId); err != nil {
		c.String(http.StatusInternalServerError, "call FavoriteRepository.DeleteByUserIdAndAppId() error:%v", err)
		return
	}

}

func (ctl FavoriteController) FindByUserIdForPage(c *gin.Context) {
	params := new(struct {
		UserId   string `form:"user_id"`
		pageNum  int    `form:"page_num"`
		PageSize int    `form:"page_size"`
	})
	if err := c.BindQuery(params); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	favorites, err := ctl.service.FindByUserIdForPage(params.UserId, params.pageNum, params.PageSize)
	if err != nil {
		c.String(http.StatusInternalServerError, "call FavoriteRepository.DeleteByUserIdAndAppId() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, &favorites)
}
