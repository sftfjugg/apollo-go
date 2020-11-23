package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/repositories"
)

type ConfigService interface {
	FindConfigByAppIdandCluster(appId, cluster, namespace, laneName string) (*models.ConfigResponse, error)
}

type configService struct {
	repository repositories.ConfigRepository
}

func NewConfigService(repository repositories.ConfigRepository) ConfigService {
	return &configService{repository: repository}
}

//这里吃进行4次查询，先查询是否有公共配置，在查询公共配置的灰度，在查询自己对应配置，最后查询自己灰度
func (s configService) FindConfigByAppIdandCluster(appId, cluster, namespace, laneName string) (*models.ConfigResponse, error) {
	m := make(map[string]string, 0)
	configResponse := new(models.ConfigResponse)
	//查询默认集群配置下的非泳道配置(公共)
	configsGlobal, err := s.repository.FindGlobalConfig(namespace, "default", "default")
	if err != nil {
		return nil, errors.Wrap(err, "find config names failed")
	}
	for i := range configsGlobal {
		config := make(map[string]string, 0)
		err := json.Unmarshal([]byte(configsGlobal[i].Configurations), &config)
		if err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal config  failed")
		}
		for k := range config {
			m[k] = config[k]
		}
	}
	//如果不是默认集群，则查找自己集群下的配置并覆盖默认集群下的配置（公共）
	if cluster != "default" {
		configsGlobal, err := s.repository.FindGlobalConfig(namespace, cluster, "default")
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsGlobal {
			config := make(map[string]string, 0)
			err := json.Unmarshal([]byte(configsGlobal[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
		}
	}
	//查询自己集群下的灰度配置（公共）
	configsLane, err := s.repository.FindGlobalConfig(namespace, cluster, laneName)
	if err != nil {
		return nil, errors.Wrap(err, "find config names failed")
	}
	for i := range configsLane {
		config := make(map[string]string, 0)
		err := json.Unmarshal([]byte(configsLane[i].Configurations), &config)
		if err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal config  failed")
		}
		for k := range config {
			m[k] = config[k]
		}
	}

	if namespace != "all" {
		//查询默认集群下的配置（非公共）
		configsDefault, err := s.repository.FindConfig(appId, "default", namespace, "default")
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsDefault {
			config := make(map[string]string, 0)
			err := json.Unmarshal([]byte(configsDefault[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
			configResponse.ReleaseKey = configsDefault[i].ReleaseKey
		}
		//查询自己集群下的配置（非公共配置）
		configsCluster, err := s.repository.FindConfig(appId, cluster, namespace, "default")
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsCluster {
			config := make(map[string]string, 0)
			err := json.Unmarshal([]byte(configsCluster[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
			configResponse.ReleaseKey = configsCluster[i].ReleaseKey
		}

		configsAll, err := s.repository.FindConfig(appId, cluster, namespace, laneName)
		if err != nil {
			return nil, errors.Wrap(err, "find config names failed")
		}
		for i := range configsAll {
			config := make(map[string]string, 0)
			err := json.Unmarshal([]byte(configsAll[i].Configurations), &config)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
			}
			for k := range config {
				m[k] = config[k]
			}
			configResponse.ReleaseKey = configsAll[i].ReleaseKey
		}
	}
	//else {
	//	configsPublic, err := s.repository.FindPublicConfig(appId)
	//	if err != nil {
	//		return nil, errors.Wrap(err, "find config names failed")
	//	}
	//	for i := range configsPublic {
	//		config := make(map[string]string, 0)
	//		err := json.Unmarshal([]byte(configsPublic[i].Configurations), &config)
	//		if err != nil {
	//			return nil, errors.Wrap(err, "json.Unmarshal config  failed")
	//		}
	//		for k := range config {
	//			m[k] = config[k]
	//		}
	//		configResponse.ReleaseKey = configsPublic[i].ReleaseKey
	//	}
	//	if cluster != "default" {
	//		configPrivates, err := s.repository.FindPrivateConfig(appId, cluster)
	//		if err != nil {
	//			return nil, errors.Wrap(err, "find config private failed")
	//		}
	//		config := make(map[string]string, 0)
	//		for i := range configPrivates {
	//			if err := json.Unmarshal([]byte(configPrivates[i].Configurations), &config); err != nil {
	//				return nil, errors.Wrap(err, "json.Unmarshal config  failed")
	//			}
	//			for k := range config {
	//				m[k] = config[k]
	//			}
	//			configResponse.ReleaseKey = configPrivates[i].ReleaseKey
	//		}
	//	}
	//}
	configResponse.Configurations = m
	configResponse.AppId = appId
	configResponse.ClusterName = cluster
	return configResponse, nil

}
