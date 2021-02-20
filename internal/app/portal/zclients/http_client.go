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
	address := single_queue.GetV()
	m, ok := address.Load(env)
	if !ok {
		return nil, errors.New("here is no adminservcie to call")
	}
	adds, ok := m.([]*models2.Address)
	if !ok {
		return nil, errors.New("here is no adminservcie to call")
	}
	i := rand.Intn(len(adds))

	//添加操作记录，操作记录需要env_id的参数和配置env_name的参数
	if env == "TEST" {
		if r.URL.RawQuery != "" {
			r.URL.RawQuery += "&env_id=1&env_name=TEST"
		} else {
			r.URL.RawQuery += "env_id=1&env_name=TEST"
		}
	} else if env == "ALIYUN" {
		if r.URL.RawQuery != "" {
			r.URL.RawQuery += "&env_id=4&env_name=ALIYUN"
		} else {
			r.URL.RawQuery += "env_id=4&env_name=ALIYUN"
		}
	} else if env == "ONLINE" {
		if r.URL.RawQuery != "" {
			r.URL.RawQuery += "&env_id=3&env_name=ONLINE"
		} else {
			r.URL.RawQuery += "env_id=3&env_name=ONLINE"
		}
	}

	r.URL.Path = url
	r.URL.Host = adds[i].InstanceId
	r.Host = adds[i].InstanceId
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
	address := single_queue.GetV()
	m, ok := address.Load(env)
	if !ok {
		return nil, errors.New("here is no adminservcie to call")
	}
	adds, ok := m.([]*models2.Address)
	if !ok {
		return nil, errors.New("here is no adminservcie to call")
	}
	i := rand.Intn(len(adds))
	url = "http://" + adds[i].InstanceId + url
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
