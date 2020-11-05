package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"net/http"
)

type ReleaseController struct {
	service services.ReleaseService
}

func NewReleaseController(service services.ReleaseService) *ReleaseController {
	return &ReleaseController{service: service}
}

func (ctl ReleaseController) Create(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	if UserID.(string) == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.Create(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ReleaseController.Create run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}

func (ctl ReleaseController) ReleaseGrayTotal(c *gin.Context) {
	env := c.Param("env")
	UserID, _ := c.Get("UserID")
	if UserID.(string) == "" {
		c.String(http.StatusUnauthorized, "UserID don't null")
		return
	}
	cookie := &http.Cookie{Name: "UserID", Value: UserID.(string), HttpOnly: true}
	c.Request.AddCookie(cookie)
	r, err := ctl.service.ReleaseGrayTotal(env, c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, "ReleaseController.ReleaseGrayTotal run failed:%v", err)
		return
	}
	c.Data(r.Code, r.ContentType, r.Data)
}
