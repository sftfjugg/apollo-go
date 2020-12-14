package services

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/models"
	"go.didapinche.com/juno-go/v2"
	"strconv"
	"strings"
)

type ConsulService interface {
	FindAddress(name string) ([]*models.Consul, error)
}

type consulService struct {
}

func NewConsulService() ConsulService {
	return &consulService{}
}

//本地consul1.1及以下版本有个bug，config不能发现部署在自己服务器的consul，并且在本地单机模式中长时间下对应服务会掉
func (s consulService) FindAddress(name string) ([]*models.Consul, error) {
	consul := juno.GetConsulRegistry()
	consulAddress := consul.ConsulAddress
	if consulAddress == "" {
		return nil, errors.New("get local ipv4 error")
	}
	c := consulapi.Config{Address: consulAddress}
	client, err := consulapi.NewClient(&c)
	if err != nil {
		return nil, errors.Wrap(err, "create consul failed")
	}
	services, _, err := client.Health().Service(name, "", false, nil)
	if err != nil {
		return nil, errors.Wrap(err, "FindAddresss failed")
	}
	if services == nil {
		return nil, errors.Wrap(err, "get 0 address")
	}
	consuls := make([]*models.Consul, 0)
	for s := range services {
		consul := new(models.Consul)
		consul.AppName = services[s].Service.Service
		ip := services[s].Service.Address

		//|| name == "apollo-plus-admin-service"本地需加
		if ip != "" || name == "apollo-plus-configservice" {
			if ip == "" {
				appJuno := juno.GetParams()
				ip = appJuno.Addr + ":" + string(appJuno.Port)
			}
			if strings.Contains(ip, ":") {
				ip = string([]byte(ip)[0:strings.Index(ip, ":")])
			}
			consul.InstanceId = ip + ":" + strconv.Itoa(services[s].Service.Port)
			consul.HomepageUrl = "http://" + ip + ":" + strconv.Itoa(services[s].Service.Port)
			consuls = append(consuls, consul)
		}
	}
	return consuls, nil
}
