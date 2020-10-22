package models

import "go.didapinche.com/time"

const AppNamespaceTableName = "AppNamespace"

//该app_namespcae只对应关联，不对应任何具体运行的项目
type AppNamespace struct {
	Id                        uint64    `gorm:"column:Id" json:"id" form:"id"`
	Name                      string    `gorm:"column:Name" json:"name" form:"name"`
	Department                string    `gorm:"column:Department" json:"department" form:"department"`
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
