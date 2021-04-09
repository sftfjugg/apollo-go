package models

import "go.didapinche.com/time"

const DingdingTableName = "Dingding"

type Dingding struct {
	Id                     uint64    `gorm:"column:Id" json:"id" form:"id"`
	DeptName               string    `gorm:"column:DeptName" json:"dept_name" form:"dept_name"`
	Name                   string    `gorm:"column:Name" json:"name" form:"name"`
	Env                    string    `gorm:"column:Env" json:"env" form:"env"`
	AppId                  string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	Type                   string    `gorm:"column:Type" json:"type" form:"type"`
	Token                  string    `gorm:"column:Token" json:"token" form:"token"`
	Level                  int       `gorm:"column:Level" json:"level" form:"level"`
	IsDeleted              bool      `gorm:"column:IsDeleted" json:"-"`
	DataChange_CreatedBy   string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_CreatedTime time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
}

func (Dingding) TableName() string {
	return DingdingTableName
}
