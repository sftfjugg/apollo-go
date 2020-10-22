package models

import "go.didapinche.com/time"

const AppTableName = "App"

//项目 1 2
type App struct {
	Id                        uint64    `gorm:"column:Id" json:"id" form:"id"`
	AppId                     string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	Name                      string    `gorm:"column:Name" json:"name" form:"name"`
	IsDeleted                 bool      `gorm:"column:IsDeleted" json:"is_deleted" form:"is_deleted"`
	DataChange_CreatedBy      string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_LastModifiedBy string    `gorm:"column:DataChange_LastModifiedBy" json:"data_change_last_modified_by" form:"data_change_last_modified_by"`
	DataChange_CreatedTime    time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
	DataChange_LastTime       time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}

func (App) TableName() string {
	return AppTableName
}
