package models

import "go.didapinche.com/time"

const ReleaseTableName = "Release"

//发布1
type Release struct {
	Id                        uint64    `gorm:"column:Id" json:"id" form:"id"`
	ReleaseKey                string    `gorm:"column:ReleaseKey" json:"release_key" form:"release_key"` //无用，留下备用
	LaneName                  string    `gorm:"column:LaneName" json:"lane_name" form:"lane_name"`       //灰度名字
	Comment                   string    `gorm:"column:Comment" json:"comment" form:"comment"`
	AppId                     string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	ClusterName               string    `gorm:"column:ClusterName" json:"cluster_name" form:"cluster_name"`
	NamespaceName             string    `gorm:"column:NamespaceName" json:"namespace_name" form:"namespace_name"`
	Configurations            string    `gorm:"column:Configurations" json:"configurations" form:"configurations"`
	IsAbandoned               bool      `gorm:"column:IsAbandoned" json:"is_abandoned" form:"is_abandoned"` //无用，保留字段
	IsDeleted                 bool      `gorm:"column:IsDeleted"`
	DataChange_CreatedBy      string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_LastModifiedBy string    `gorm:"column:DataChange_LastModifiedBy" json:"data_change_last_modified_by" form:"data_change_last_modified_by"`
	DataChange_CreatedTime    time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
	DataChange_LastTime       time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}

func (Release) TableName() string {
	return ReleaseTableName
}

//INSERT INTO `Release` (`Id`,`ReleaseKey`,`Comment`,`AppId`,`ClusterName`,`NamespaceName`,`Configurations`,`IsAbandoned`,`IsDeleted`) VALUES (0,'2020-10-23 18:40:02','xbGPTfgdcCpntzbVrHaLNOiMaXpuFJ','1000044581','default','test','{"myName":"123434"}',0,0)
