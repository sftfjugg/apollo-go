package models

import "go.didapinche.com/time"

const HistoryTableName = "History"

//项目 1 2
type History struct {
	Id                     uint64    `gorm:"column:Id" json:"id" form:"id"`
	AppId                  string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	UserId                 string    `gorm:"column:UserId" json:"user_id" form:"user_id"`
	UserName               string    `gorm:"column:UserName" json:"user_name" form:"user_name"`
	IsDeleted              bool      `gorm:"column:IsDeleted" json:"-"`
	DataChange_CreatedTime time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
}

func (History) TableName() string {
	return HistoryTableName
}
