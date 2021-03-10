package models

type Consul struct {
	AppName     string `json:"appName"`
	InstanceId  string `json:"instanceId"`
	HomepageUrl string `json:"homepageUrl"`
}
