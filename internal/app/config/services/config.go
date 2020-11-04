package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/repositories"
)

type ConfigService interface {
	FindConfigByAppIdandCluster(appId, cluster, namespace string) (*models.ConfigResponse, error)
}

type configService struct {
	repository repositories.ConfigRepository
}

func NewConfigService(repository repositories.ConfigRepository) ConfigService {
	return &configService{repository: repository}
}

func (s configService) FindConfigByAppIdandCluster(appId, cluster, namespace string) (*models.ConfigResponse, error) {
	m := make(map[string]string)
	configResponse := new(models.ConfigResponse)
	//查询公共全局配置
	configsGlobal, err := s.repository.FindGlobalConfig("default")
	if err != nil {
		return nil, errors.Wrap(err, "find config names failed")
	}
	for i := range configsGlobal {
		config := make(map[string]string)
		err := json.Unmarshal([]byte(configsGlobal[i].Configurations), &config)
		if err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal config  failed")
		}
		for k := range config {
			m[k] = config[k]
		}
	}
	if cluster != "default" {
		configsGlobal, err := s.repository.FindGlobalConfig(cluster)
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsGlobal {
			config := make(map[string]string)
			err := json.Unmarshal([]byte(configsGlobal[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
		}
	}

	if namespace != "all" {
		configsDefault, err := s.repository.FindConfig(appId, "default", namespace)
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsDefault {
			config := make(map[string]string)
			err := json.Unmarshal([]byte(configsDefault[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
			configResponse.ReleaseKey = configsDefault[i].ReleaseKey
		}
		configsAll, err := s.repository.FindConfig(appId, cluster, namespace)
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsAll {
			config := make(map[string]string)
			err := json.Unmarshal([]byte(configsAll[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
			configResponse.ReleaseKey = configsAll[i].ReleaseKey
		}
	} else {
		configsPublic, err := s.repository.FindPublicConfig(appId)
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsPublic {
			config := make(map[string]string)
			err := json.Unmarshal([]byte(configsPublic[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
			configResponse.ReleaseKey = configsPublic[i].ReleaseKey
		}
		if cluster != "default" {
			configPrivates, err := s.repository.FindPrivateConfig(appId, cluster)
			if err != nil {
				return nil, errors.Wrap(err, "find config private failed")
			}
			config := make(map[string]string)
			for i := range configPrivates {
				if err := json.Unmarshal([]byte(configPrivates[i].Configurations), &config); err != nil {
					return nil, errors.Wrap(err, "json.Unmarshal config  failed")
				}
				for k := range config {
					m[k] = config[k]
				}
				configResponse.ReleaseKey = configPrivates[i].ReleaseKey
			}
		}
	}
	configResponse.Configurations = m
	configResponse.AppId = appId
	configResponse.ClusterName = cluster
	return configResponse, nil

}
