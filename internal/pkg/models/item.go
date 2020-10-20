package models

import "go.didapinche.com/time"

const ItemTableName = "Item"

//配置文件 1
type Item struct {
	Id                        uint64    `gorm:"column:Id" json:"id" form:"id"`
	NamespaceId               uint64    `gorm:"column:NamespaceId" json:"namespace_id" form:"namespace_id"`
	Key                       string    `gorm:"column:Key" json:"key" form:"value"`
	Value                     string    `gorm:"column:Value" json:"value" form:"value"`
	Status                    uint64    `gorm:"column:Status" json:"status" from:"status"`
	Comment                   string    `gorm:"column:Comment" json:"comment" form:"comment"`
	Describe                  string    `gorm:"column:Describe" json:"describe" form:"describe"`
	IsDeleted                 bool      `gorm:"column:IsDeleted" json:"-" form:"is_deleted"`
	DataChange_CreatedBy      string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_LastModifiedBy string    `gorm:"column:DataChange_LastModifiedBy" json:"data_change_last_modified_by" form:"data_change_last_modified_by"`
	DataChange_CreatedTime    time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
	DataChange_LastTime       time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}

func (Item) TableName() string {
	return ItemTableName
}
