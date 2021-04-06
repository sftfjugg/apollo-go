package services

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
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
}

func NewReleaseService(httpClient *zclients.HttpClient) ReleaseService {
	return &releaseService{httpClient: httpClient}
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

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	response, err := s.httpClient.HttpDo("/release", env, c.Request)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response.Code == 200 {

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
