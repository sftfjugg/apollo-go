package models

import "go.didapinche.com/time"

const ReleaseHistoryTableName = "ReleaseHistory"

//发布历史 1
type ReleaseHistory struct {
	Id                        uint64    `gorm:"column:Id" json:"id" form:"id"`
	AppId                     string    `gorm:"column:AppId" json:"app_id" form:"app_id"`
	ClusterName               string    `gorm:"column:ClusterName" json:"cluster_name" form:"cluster_name"`
	NamespaceName             string    `gorm:"column:NamespaceName" json:"namespace_name" form:"namespace_name"`
	BranchName                string    `gorm:"column:BranchName" json:"branch_name" form:"branch_name"`
	ReleaseId                 uint64    `gorm:"column:ReleaseId" json:"release_id" form:"release_id"`
	PreviousReleaseId         uint64    `gorm:"column:PreviousReleaseId" json:"previous_release_id" form:"previous_release_id"`
	Operation                 string    `gorm:"column:Operation" json:"operation" form:"operation"`
	OperationContext          string    `gorm:"column:OperationContext" json:"operation_context" form:"operation_context"`
	ReleaseContext            string    `gorm:"column:ReleaseContext" json:"release_context" form:"release_context"`
	IsDeleted                 bool      `gorm:"column:IsDeleted" json:"-" form:"is_deleted"`
	DataChange_CreatedBy      string    `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	DataChange_LastModifiedBy string    `gorm:"column:DataChange_LastModifiedBy" json:"data_change_last_modified_by" form:"data_change_last_modified_by"`
	DataChange_CreatedTime    time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
	DataChange_LastTime       time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}

func (ReleaseHistory) TableName() string {
	return ReleaseHistoryTableName
}
