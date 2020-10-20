package models

const ClusterTableName = "Cluster"

//集群配置 1
type Cluster struct {
	Base            Base   `gorm:"embedded"`
	Name            string `gorm:"column:Name" json:"name" form:"name"`
	AppId           string `gorm:"column:AppId" json:"app_id" form:"name"`
	ParentClusterId uint64 `gorm:"column:ParentClusterId" json:"parent_cluster_id" form:"parent_cluster_id"`
}

func (Cluster) TableName() string {
	return ClusterTableName
}
