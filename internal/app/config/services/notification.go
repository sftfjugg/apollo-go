package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/single_queue"
	"time"
)

type NotificationMessageService interface {
	CompareV(appid, cluster, notifications, lane string) ([]*models.Notification, error)
}

type notificationMessageService struct {
}

func NewNotificationMessageService() NotificationMessageService {
	return &notificationMessageService{}
}

//监视配置文件是否改变
func (s notificationMessageService) CompareV(appid, cluster, notifications, lane string) ([]*models.Notification, error) {
	tempMap := make([]map[string]interface{}, 0)
	params := make([]*models.Notification, 0) //返回值为了匹配客户端使用数组，但是数组中只有一组数据
	err := json.Unmarshal([]byte(notifications), &tempMap)
	if err != nil {
		return nil, errors.Wrap(err, "json bind failed")
	}

	mav := make(map[string]models.Version)
	for i := range tempMap {
		namespaceName := tempMap[i]["namespaceName"].(string)
		v := models.Version{Max: tempMap[i]["notificationId"].(float64), Index: i}
		//这里添加6个监视，分别为默认配置，公共默认配置，如果存在非默认集群，添加非默认集群配置，非默认集群灰度配置，存在灰度则在添加2个
		mav[appid+"+"+cluster+"+"+namespaceName] = v
		mav["public_global_config"+"+"+cluster+"+"+namespaceName] = v
		if cluster != "default" {
			mav[appid+"+"+"default"+"+"+namespaceName] = v
			mav["public_global_config"+"+"+"default"+"+"+namespaceName] = v
		}
		if lane != "default" {
			mav[appid+lane+"+"+"default"+"+"+namespaceName] = v
			mav["public_global_config"+lane+"+"+cluster+"+"+namespaceName] = v
		}

	}

	m := single_queue.GetV()
	i := 0
	typ := false //标示位置，true则说明配置项动态改变了
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
loop:
	for range ticker.C {
		for k, v := range mav {
			if float64(m[k]) > v.Max {
				tempMap[v.Index]["notificationId"] = float64(m[k])
				typ = true
				break loop
			}
		}
		i++
		if i >= 60 {
			break loop
		}
	}

	for i := range tempMap {
		notification := new(models.Notification)
		notification.NamespaceName = tempMap[i]["namespaceName"].(string)
		notification.NotificationId = int(tempMap[i]["notificationId"].(float64))
		message := new(models.Messages)
		message.Details = make(map[string]int)
		message.Details[notification.NamespaceName] = notification.NotificationId
		notification.Messages = message
		params = append(params, notification)
	}

	if typ {
		return params, nil
	}

	return nil, nil
}
