package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/data"
)

type DateController struct {
}

func (d DateController) ImportDate(c *gin.Context) {
	data.ImportData()
	c.String(200, "倒数据结束")
}
