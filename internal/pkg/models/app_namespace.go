package models

import "go.didapinche.com/time"

const AppNamespaceTableName = "AppNamespace"

//配置文件， Format为对应集群名称 1 2
type AppNamespace struct {
	Id                        uint64    `gorm:"column:Id" json:"id" form:"id"`
	Name                      string    `gorm:"column:Name" json:"name" form:"name"` //全局唯一
	AppId                     string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	ClusterName               string    `gorm:"column:ClusterName" json:"cluster_name" form:"cluster_name"`
	Format                    string    `gorm:"column:Format" json:"format" form:"format"` //类型，目前只支持properties
	IsPublic                  bool      `gorm:"column:IsPublic" json:"is_public" form:"is_public"`
	Comment                   string    `gorm:"column:Comment" json:"comment" form:"comment"`
	IsDeleted                 bool      `gorm:"column:IsDeleted" json:"-"`
	DataChange_CreatedBy      string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_LastModifiedBy string    `gorm:"column:DataChange_LastModifiedBy" json:"data_change_last_modified_by" form:"data_change_last_modified_by"`
	DataChange_CreatedTime    time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
	DataChange_LastTime       time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}

func (AppNamespace) TableName() string {
	return AppNamespaceTableName
}