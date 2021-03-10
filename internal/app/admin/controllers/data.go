package controllers

import (
	"github.com/gin-gonic/gin"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/data"
	"net/http"
)

type DateController struct {
}

func NewDateController() *DateController {
	return &DateController{}
}

func (d DateController) Health(c *gin.Context) {
	c.String(200, "检查成功")
}

//数据倒入
func (d DateController) ImportDate(c *gin.Context) {
	param := new(struct {
		DB1 string `json:"db1"`
		DB2 string `json:"db2"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	data.ImportData(param.DB1, param.DB2)
	c.String(200, "倒数据结束")
}

//数据倒入
func (d DateController) UpdateDate(c *gin.Context) {
	param := new(struct {
		DB1 string `json:"db1"`
		DB2 string `json:"db2"`
	})
	if err := c.ShouldBind(param); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	data.UpadteDate(param.DB2)
	c.String(200, "倒数据结束")
}
