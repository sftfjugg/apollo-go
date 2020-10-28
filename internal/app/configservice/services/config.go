package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/configservice/repositories"
)

type ConfigService interface {
	FindConfigByAppIdandCluster(appId, cluster string) (*models.ConfigResponse, error)
}

type configService struct {
	repository repositories.ConfigRepository
}

func NewConfigService(repository repositories.ConfigRepository) ConfigService {
	return &configService{repository: repository}
}

func (s configService) FindConfigByAppIdandCluster(appId, cluster string) (*models.ConfigResponse, error) {
	names, err := s.repository.FindPublicConfigName(appId)
	if err != nil {
		return nil, errors.Wrap(err, "find config names failed")
	}
	m := make(map[string]string)
	for i := range names {
		configsPublic, err := s.repository.FindPublicConfig(appId, names[i].Name)
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for j := range configsPublic {
			config := make(map[string]string)
			err := json.Unmarshal([]byte(configsPublic[j].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
		}
	}
	configPrivates, err := s.repository.FindPrivateConfig(appId, cluster)
	if err != nil {
		return nil, errors.Wrap(err, "find config private failed")
	}
	config := make(map[string]string)
	configResponse := new(models.ConfigResponse)
	for i := range configPrivates {
		if err := json.Unmarshal([]byte(configPrivates[i].Configurations), &config); err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal config  failed")
		}
		for k := range config {
			m[k] = config[k]
		}
		configResponse.ReleaseKey = configPrivates[i].ReleaseKey
	}
	configResponse.Configurations = m
	configResponse.AppId = appId
	configResponse.ClusterName = cluster
	return configResponse, nil

}
