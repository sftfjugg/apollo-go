package models

const ServerConfigTableName = "ServerConfig"

//服务器自身配置 12
type ServerConfig struct {
	Base    Base   `gorm:"embedded"`
	Key     string `gorm:"column:Key"  json:"key" form:"value"`
	Value   string `gorm:"column:Value"  json:"value" form:"value"`
	Comment string `gorm:"column:Comment"  json:"comment" form:"comment"`
	Cluster string `gorm:"column:Cluster"  json:"cluster" form:"cluster"`
}

func (ServerConfig) TableName() string {
	return ServerConfigTableName
}
