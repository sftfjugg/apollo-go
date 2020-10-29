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
	notification := new(models.Notification)
	params := make([]*models.Notification, 0) //返回值为了匹配客户端使用数组，但是数组中只有一组数据
	err := json.Unmarshal([]byte(notifications), &tempMap)
	if err != nil {
		return nil, errors.Wrap(err, "json bind failed")
	}
	namespaceName := tempMap[0]["namespaceName"].(string)
	max := tempMap[0]["notificationId"].(float64)
	key := appid + "+" + cluster + "+" + namespaceName
	m := single_queue.GetV()
	i := 0
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		if float64(m[key]) > max {
			break
		}
		i++
		if i >= 60 {
			break
		}
	}
	notification.NamespaceName = namespaceName
	notification.NotificationId = int(m[key])
	message := new(models.Messages)
	message.Details = make(map[string]int)
	message.Details[key] = int(m[key])
	notification.Messages = message
	params = append(params, notification)
	if float64(m[key]) > max {
		return params, nil
	}
	return nil, nil
}
