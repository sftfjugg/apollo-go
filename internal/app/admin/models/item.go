package models

import (
	"go.didapinche.com/time"
)

type ItemPage struct {
	Total int     `json:"total"`
	Items []*Item `json:"items"`
}

type Item struct {
	Id                        uint64    `gorm:"column:Id" json:"id" form:"id"`
	Name                      string    `gorm:"column:Name" json:"name" form:"name"`
	AppId                     string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	AppName                   string    `gorm:"column:AppName" json:"app_name" form:"app_name"`
	ClusterName               string    `gorm:"column:ClusterName" json:"cluster_name" form:"cluster_name"` //灰度使用
	LaneName                  string    `gorm:"column:LaneName" json:"lane_name" form:"lane_name"`
	NamespaceId               uint64    `gorm:"column:NamespaceId" json:"namespace_id" form:"namespace_id"`
	Key                       string    `gorm:"column:Key" json:"key" form:"value"`
	Value                     string    `gorm:"column:Value" json:"value" form:"value"`
	ReleaseValue              string    `gorm:"column:ReleaseValue" json:"release_value" form:"release_value"`
	Status                    uint64    `gorm:"column:Status" json:"status" from:"status"`
	Format                    string    `gorm:"column:Format" json:"format" form:"format"` //类型
	Comment                   string    `gorm:"column:Comment" json:"comment" form:"comment"`
	Describe                  string    `gorm:"column:Describe" json:"describe" form:"describe"`
	DataChange_CreatedBy      string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_LastModifiedBy string    `gorm:"column:DataChange_LastModifiedBy" json:"data_change_last_modified_by" form:"data_change_last_modified_by"`
	DataChange_CreatedTime    time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
	DataChange_LastTime       time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}

type Count struct {
	Count int `form:"count"`
}

//Select I.Id,I.Key,I.Value,I.NamespaceId,A.Name,A.AppId,A.AppName,A.ClusterName,A.LaneName,I.Status,I.Comment,I.Describe,I.DataChange_CreatedBy,I.DataChange_LastModifiedBy,I.DataChange_CreatedTime,I.DataChange_LastTime     from AppNamespace A,Item I where I.Key like "%m%" and A.Id=I.NamespaceId;
