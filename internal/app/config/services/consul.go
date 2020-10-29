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
	services, err := client.Agent().Services()
	if err != nil {
		return nil, errors.Wrap(err, "FindAddresss failed")
	}
	if services == nil {
		return nil, errors.Wrap(err, "get 0 address")
	}
	consuls := make([]*models.Consul, 0)
	for s := range services {
		if strings.EqualFold(services[s].Service, name) {
			consul := new(models.Consul)
			consul.AppName = services[s].Service
			ip := services[s].Address
			if ip == "" {
				appJuno := juno.GetParams()
				ip = appJuno.Addr + ":" + string(appJuno.Port)
			}
			if strings.Contains(ip, ":") {
				ip = string([]byte(ip)[0:strings.Index(ip, ":")])
			}
			consul.InstanceId = ip + ":" + strconv.Itoa(services[s].Port)
			consul.HomepageUrl = "http://" + ip + ":" + strconv.Itoa(services[s].Port)
			consuls = append(consuls, consul)
		}
	}
	return consuls, nil
}
