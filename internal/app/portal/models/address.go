package models

type Address struct {
	AppName     string `json:"appName" form:"appName"`
	InstanceId  string `json:"instanceId" form:"instanceId"`
	HomepageUrl string `json:"homepageUrl" form:"homepageUrl"`
}
