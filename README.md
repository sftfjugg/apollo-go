# apollo+
[toc]
## 项目背景
Apollo+是使用golang开发的配置中心，在市场的配置中心基础上，增加了泳道，公共配置，流水线集成，批量管理，集成limos，钉钉报警等功能等功能
## 架构介绍
Apollo+包含3个子项目，分别上admin，config，portal，portal是管理端，只需要在线上部署一个，config是与客户端连接的一个轻量级web服务，具有最重要的获取配置和动态更新接口，需要在每个环境中部署，admin上操作当前环境配置的web服务，需要在每个环境中至少部署一次
### 维护项目
apollo+需要维护apollo+项目，agollo（golang客户端）和对携程apollo的客户端进行改动
## 源码讲解
### portal
portal相对简单，主要做权限管理和流量转发，将制定环境的流量转发给制定环境的admin
### admin
admin进行了大量的sql操作，主要上对当前环境的配置进行增删改查等操作
### config
config只连客户端，主要是4个接口
#### /services/config
获取当前环境config集群ip列表，主要通过consul获取
```
func (ctl ConsulController) FindConfigService(c *gin.Context) {
	consul, err := ctl.services.FindAddress("apollo-plus-configservice")
	if err != nil {
		c.String(http.StatusBadRequest, "call ConsulService.FindConsulByName error:%v", err)
		return
	}
	c.JSON(http.StatusOK, consul)
}
```
主要使用的consul的api
```

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

		//|| name == "apollo-plus-admin-service"本地如果consul不健康需加
		if ip != "" || name == "apollo-plus-configservice" || name == "apollo-plus-admin-service" {
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

```

#### /services/admin
获得当前环境的admin，原理与config相同
#### /configs/:appId/:clusterName/:namespace
这个是客户端用来获取配置的接口
在apollo中，公共配置namespaces名字是public_global_config，因此，获取配置需要进行6次查询进行配置覆盖

```
//这里吃进行6次查询，先查询是否有公共配置，在查询公共配置的灰度，在查询自己对应配置，最后查询自己灰度
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
	configResponse.Configurations = m
	configResponse.AppId = appId
	configResponse.ClusterName = cluster
	return configResponse, nil

}
```

#### /notifications/v2
改接口是查询配置更新使用，内部原理是在请求到来后不相应请求，进行60秒的版本号比对，如果版本号在60秒内发生变更，则立刻返回200，如果没有变更则返回304客户端重新拉去配置，会有一个协程一直查询最新sql获得版本号并存入一个全部单列map结构中

## 部署方案
建议线上部署3台以上config，其余环境可部署2台以上，admin每个环境都部署2台，portal线上部署一台
## 其他系统集成

