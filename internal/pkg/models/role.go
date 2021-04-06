package models

import "go.didapinche.com/time"

const RoleTableName = "Role"

//项目 1 2
type Role struct {
	Id                     uint64    `gorm:"column:Id" json:"id" form:"id"`
	AppId                  string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	Namespace              string    `gorm:"column:Namespace" json:"namespace" form:"namespace"`
	Cluster                string    `gorm:"column:cluster" json:"cluster" form:"cluster"`
	Env                    string    `gorm:"column:Env" json:"env" form:"env"`
	UserID                 string    `gorm:"column:UserId" json:"user_id" form:"user_id"`
	UserName               string    `gorm:"column:UserName" json:"user_name" form:"user_name"` //部门id
	Level                  int       `gorm:"column:Level" json:"level" form:"level"`
	IsDeleted              bool      `gorm:"column:IsDeleted" json:"-"`
	DataChange_CreatedBy   string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_CreatedTime time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
}

func (Role) TableName() string {
	return AppTableName
}
