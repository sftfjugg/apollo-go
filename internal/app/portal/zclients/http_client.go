package zclients

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/single_queue"
	"io/ioutil"
	"math/rand"
	"net/http"
)

//http client方法封装，用于远程调用adminservice
type HttpClient struct {
	client *http.Client
}

func NewHttpClient(client *http.Client) *HttpClient {
	return &HttpClient{client: client}
}

//http方法封装,拿到gin的request，去除r的RequseURI和RemoteAddr，通过env选择对应环境的url，修改URL地址，获得请求
func (s HttpClient) HttpDo(url, env string, r *http.Request) (*models2.Response, error) {
	m := single_queue.GetV()
	if len(m[env]) == 0 {
		return nil, errors.New("There is no adminservcie to call")
	}
	i := rand.Intn(len(m[env]))
	r.URL.Path = url
	r.URL.Host = m[env][i].InstanceId
	r.Host = m[env][i].InstanceId
	r.URL.Scheme = "http"
	r.RequestURI = ""
	r.RemoteAddr = ""
	res, err := s.client.Do(r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpDo request failed")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "HttpDo request failed")
	}
	response := new(models2.Response)
	response.Data = body
	response.Code = res.StatusCode
	response.ContentType = res.Header.Get("Content-Type")
	return response, nil
}

//http方法封装,拿到gin的request，去除r的RequseURI和RemoteAddr，通过env选择对应环境的url，修改URL地址，获得请求
func (s HttpClient) HttpPost(url, env string, data interface{}) (*models2.Response, error) {
	m := single_queue.GetV()
	if len(m[env]) == 0 {
		return nil, errors.New("There is no adminservcie to call")
	}
	i := rand.Intn(len(m[env]))
	url = "http://" + m[env][i].InstanceId + url
	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "json format error:")
	}
	r := bytes.NewBuffer(b)
	res, err := s.client.Post(url, "application/json;charset=UTF-8", r)
	if err != nil {
		return nil, errors.Wrap(err, "json format error:")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll(res.Body) failed")
	}
	response := new(models2.Response)
	response.Data = body
	response.Code = res.StatusCode
	response.ContentType = res.Header.Get("Content-Type")
	return response, nil
}
