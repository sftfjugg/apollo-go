package models

//数据库响应模型
type Config struct {
	AppId          string `gorm:"column:AppId"`
	ClusterName    string `gorm:"column:ClusterName"`
	NamespaceName  string `gorm:"column:NamespaceName"`
	Configurations string `gorm:"column:Configurations"`
	ReleaseKey     string `gorm:"column:ReleaseKey"`
}

//返回值
type ConfigResponse struct {
	AppId          string            `json:"appId"`
	ClusterName    string            `json:"cluster"`
	NamespaceName  string            `json:"namespaceName"`
	Configurations map[string]string `json:"configurations"`
	ReleaseKey     string            `json:"releaseKey"`
}
