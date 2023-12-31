package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/uber/tchannel-go"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/zclients"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/dingding"
	"go.didapinche.com/goapi/plat_limos_rpc"
	"go.didapinche.com/goapi/uic_service_api"
	"go.didapinche.com/goapi/user_department_service_api"
	"go.didapinche.com/time"
	"io/ioutil"
	"net/http"
	"strings"
	time2 "time"
)

type ReleaseService interface {
	Create(env, userID string, c *gin.Context) (*models2.Response, error)
	Creates(env, userID string, c *gin.Context) (*models2.Response, error)
	ReleaseGrayTotal(env string, r *http.Request) (*models2.Response, error)
}

type releaseService struct {
	httpClient         *zclients.HttpClient
	dingding           dingding.DingDing
	dingdingRepository repositories.DingdingRepository
	uic                uic_service_api.TChanUicService
	department         user_department_service_api.TChanDepartmentService
	limos              plat_limos_rpc.TChanLimosService
	user               user_department_service_api.TChanUserService
}

func NewReleaseService(httpClient *zclients.HttpClient,
	dingding dingding.DingDing,
	department user_department_service_api.TChanDepartmentService,
	uic uic_service_api.TChanUicService,
	dingdingRepository repositories.DingdingRepository,
	limos plat_limos_rpc.TChanLimosService,
	user user_department_service_api.TChanUserService) ReleaseService {
	return &releaseService{httpClient: httpClient,
		dingding:           dingding,
		uic:                uic,
		department:         department,
		limos:              limos,
		dingdingRepository: dingdingRepository,
		user:               user}
}

//封装原writer 记录数据
type bodyLogWriter struct {
	gin.ResponseWriter
	c *gin.Context
}

func (s releaseService) Creates(env, userID string, c *gin.Context) (*models2.Response, error) {
	blw := &bodyLogWriter{
		ResponseWriter: c.Writer,
		c:              c,
	}
	//记录响应数据
	c.Writer = blw
	//记录请求的json数据

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	releases := make([]*models.ReleaseRequest, 0)
	if err := json.Unmarshal(bodyBytes, &releases); err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal run failed")
	}

	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll run failed")
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	response, err := s.httpClient.HttpDo("/releases", env, c.Request)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response.Code == 200 {
		go func() {
			for _, r := range releases {
				s.sendDingding(env, userID, r)
			}
		}()
	}
	return response, nil
}

func (s releaseService) Create(env, userID string, c *gin.Context) (*models2.Response, error) {
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

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	response, err := s.httpClient.HttpDo("/release", env, c.Request)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	if response.Code == 200 {
		go func() {
			s.sendDingding(env, userID, release)
		}()
	}
	return response, nil
}

func (s releaseService) sendDingding(env, userID string, release *models.ReleaseRequest) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	text := ""
	title := ""
	level := 1
	owners := make(map[string]int)
	userIDs := make([]string, 0)
	tp := "应用:"
	if release.AppId != "public_global_config" {
		title += release.AppId
		tp = "应用"
		text += "# Apollo应用配置变动通知  \n  应用:"
		text += "[" + release.AppId + "](http://pass.didapinche.com/apollo/application/list?cluster=" + release.ClusterName + "&app_name=" + release.AppId + "&env=" + env + ")  \n"
	} else {
		title += "公共配置"
		tp = "公共配置"
		text += "# Apollo公共配置变动通知 \n"
	}
	text += " 环境:   *" + env + "*  \n"
	text += " 集群:   *" + release.ClusterName + "*  \n"
	text += " 命名空间:   *" + release.Name + "*  \n"
	text += "  \n 操作人: *" + userID + "*"
	text += "  \n 操作时间: *" + time.Now().String() + "*    \n"
	if release.AppId == "public_global_config" {
		text += "*[查看详情](http://pass.didapinche.com/apollo/public/list?env=" + env + "&cluster=" + release.ClusterName + "&app_name=public_global_config)*"
	}
	text += "\n ---  \n  "
	for i, k := range release.Keys {
		text += "***" + k + " = " + release.Values[i] + "***  \n"
	}
	text += "\n ---  \n  "
	ctx, _ := tchannel.NewContextBuilder(time2.Second).Build()
	node, err := s.department.FindAllNodes(ctx)
	if err == nil {
		m := s.getMap(node)

		if release.DeptName == "" {
			ctx2, _ := tchannel.NewContextBuilder(time2.Second).Build()
			app, err := s.limos.FindAppForPage(ctx2, release.AppId, "", "", "", "", 0, "all", 1000, 1, "", 0)
			if err == nil {
				for _, a := range app.Apps {
					if release.AppId == a.Name {
						release.DeptName = a.DevGroupName
						userIDs = append(userIDs, a.OwnerIds...)
						level = int(a.Level)
					}
				}
			}
		}
		deptnames := strings.Split(release.DeptName, ",")
		tokens := make(map[string]int, 0)
		for _, d := range deptnames {
			token := ""
			for {
				if token != "" {
					break
				}
				ding, err := s.dingdingRepository.Find(tp, d, env, level)
				if err == nil {
					token = ding.Token
					if token != "" {
						tokens[token] = 1
					}
					d = m[d]
					if d == "" {
						break
					}
				} else {
					break
				}
			}
		}

		for _, user := range userIDs {
			ctx, _ := tchannel.NewContextBuilder(time2.Second).Build()
			owner, err := s.user.FindByUid(ctx, user)
			if err == nil {
				owners[owner.Phone] = 1
			}
		}
		phones := make([]string, 0)
		for k, _ := range owners {
			text += "[@" + k + "]() "
			phones = append(phones, k)
		}

		msg := &dingding.DingMessage{
			MessageType: "markdown",
			Markdown: dingding.Markdown{
				Title: title,
				Text:  text,
			},
			At: dingding.At{
				AtMobiles: phones,
				IsAtAll:   false},
		}

		for t, _ := range tokens {
			s.dingding.SendMessage(t, msg)
		}
	}
}

func (s releaseService) ReleaseGrayTotal(env string, r *http.Request) (*models2.Response, error) {
	response, err := s.httpClient.HttpDo("/release_gray_total", env, r)
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient HttpDo run failed")
	}
	return response, nil
}

//获得子对应父的层级关系
func (s releaseService) getMap(node *user_department_service_api.DepartmentTreeNode) map[string]string {
	tmpMap := make(map[string]*user_department_service_api.DepartmentTreeNode, 0)
	result := make(map[string]string)
	//每层节点
	tmpSlice := make([]string, 0)

	tmpSlice = append(tmpSlice, node.Name)
	tmpMap[node.Name] = node
	result[node.Name] = ""

	for len(tmpSlice) != 0 {
		tmp := make([]string, 0)
		for _, tmpNodeID := range tmpSlice {
			nodes := tmpMap[tmpNodeID].ChildrenNode
			if nil != nodes && len(nodes) > 0 {
				for _, tmpNode := range nodes {
					tmp = append(tmp, tmpNode.Name)
					tmpMap[tmpNode.Name] = tmpNode
					result[tmpNode.Name] = tmpNodeID
				}
			}
		}
		tmpSlice = tmp
	}

	return result
}
