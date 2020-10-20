package services

import (
	"apollo-adminserivce/internal/app/configservice/models"
	"apollo-adminserivce/internal/pkg/utils/netutil"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
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
	host := netutil.GetLocalIP4()
	if host == "" {
		return nil, errors.New("get local ipv4 error")
	}
	addreee := host + ":7888"
	c := consulapi.Config{Address: addreee}
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
			if services[s].Address == "" {
				services[s].Address = host
			}
			consul.InstanceId = services[s].Address + ":" + strconv.Itoa(services[s].Port)
			consul.HomepageUrl = "http://" + services[s].Address + ":" + strconv.Itoa(services[s].Port)
			consuls = append(consuls, consul)
		}
	}
	return consuls, nil
}
