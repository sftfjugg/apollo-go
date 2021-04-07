package services

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/dingding"
	"go.didapinche.com/time"
	"io/ioutil"
	"net/http"
)

type ReleaseService interface {
	Create(env string, c *gin.Context) (*models2.Response, error)
	Creates(env string, c *gin.Context) (*models2.Response, error)
	ReleaseGrayTotal(env string, r *http.Request) (*models2.Response, error)
}

type releaseService struct {
	httpClient *zclients.HttpClient
	dingding   dingding.DingDing
}

func NewReleaseService(httpClient *zclients.HttpClient, dingding dingding.DingDing) ReleaseService {
	return &releaseService{httpClient: httpClient, dingding: dingding}
}

//封装原writer 记录数据
type bodyLogWriter struct {
	gin.ResponseWriter
	c *gin.Context
}

func (s releaseService) Creates(env string, c *gin.Context) (*models2.Response, error) {
	blw := &bodyLogWriter{
		ResponseWriter: c.Writer,
		c:              c,
	}
	//记录响应数据
	c.Writer = blw
	//记录请求的json数据
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll run failed")
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	response, err := s.httpClient.HttpDo("/releases", env, c.Request)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response.Code == 200 {

	}
	return response, nil
}

func (s releaseService) Create(env string, c *gin.Context) (*models2.Response, error) {
	blw := &bodyLogWriter{
		ResponseWriter: c.Writer,
		c:              c,
	}
	//记录响应数据
	c.Writer = blw
	//记录请求的json数据
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll run failed")
	}
	release := new(models.ReleaseRequest)
	if err := json.Unmarshal(bodyBytes, release); err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal run failed")
	}

	text := ""
	if release.AppId != "public_global_config" {
		text += "### Apollo应用配置变动通知  \n 应用  \n"
		text += "* [" + release.AppId + "](http://pass.didapinche.com/apollo/application/list?cluster=" + release.ClusterName + "&app_name=" + release.AppId + "&env=" + env + ")"
	} else {
		text += "### Apollo公共配置变动通知"
	}
	text += "  \n环境  \n*" + env
	text += "  \n集群  \n*" + release.ClusterName
	text += "  \n命名空间  \n*" + release.Name
	text += "  \n{修改}配置  \n"

	for _, k := range release.Keys {
		text += "*" + k + ":" + "test" + "  \n"
	}

	text += "  \n操作人:" + release.Operator
	text += "  \n操作时间:" + time.Now().String()
	msg := &dingding.DingMessage{
		MessageType: "markdown",
		Markdown: dingding.Markdown{
			Title: "应用变更配置",
			Text:  text,
		},
		At: dingding.At{IsAtAll: false},
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	response, err := s.httpClient.HttpDo("/release", env, c.Request)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response.Code == 200 {
		s.dingding.SendMessage("676ef385699499e977cdc4db3609b13fc1098ae04c847f3f5285d428c0cd0497", msg)
	}
	return response, nil
}

func (s releaseService) ReleaseGrayTotal(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/release_gray_total", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}
