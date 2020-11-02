package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/single_queue"
	"time"
)

type NotificationMessageService interface {
	CompareV(appid, cluster, notifications string) ([]*models.Notification, error)
}

type notificationMessageService struct {
}

func NewNotificationMessageService() NotificationMessageService {
	return &notificationMessageService{}
}

//监视配置文件是否改变
func (s notificationMessageService) CompareV(appid, cluster, notifications string) ([]*models.Notification, error) {
	tempMap := make([]map[string]interface{}, 0)
	params := make([]*models.Notification, 0) //返回值为了匹配客户端使用数组，但是数组中只有一组数据
	err := json.Unmarshal([]byte(notifications), &tempMap)
	if err != nil {
		return nil, errors.Wrap(err, "json bind failed")
	}
	max := make([]float64, 0)
	key := make([]string, 0)
	for i := range tempMap {
		namespaceName := tempMap[i]["namespaceName"].(string)
		max = append(max, tempMap[i]["notificationId"].(float64))
		key = append(key, appid+"+"+cluster+"+"+namespaceName)
	}

	m := single_queue.GetV()
	i := 0
	typ := false //标示位置，true则说明配置项动态改变了
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
loop:
	for range ticker.C {
		for i := range tempMap {
			k := key[i]
			if float64(m[k]) > max[i] {
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
		notification.NotificationId = int(m[appid+"+"+cluster+"+"+notification.NamespaceName])
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
