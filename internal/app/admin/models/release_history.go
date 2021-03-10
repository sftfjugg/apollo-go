package models

import (
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type ReleaseHistoryPage struct {
	ReleaseHistory []*ReleaseHistory `json:"release_history"`
	Total          int               `json:"total"`
}

//发布历史 1
type ReleaseHistory struct {
	Id                   uint64         `gorm:"column:Id" json:"id" form:"id"`
	AppId                string         `gorm:"column:AppId" json:"app_id" form:"app_id"`
	ClusterName          string         `gorm:"column:ClusterName" json:"cluster_name" form:"cluster_name"`
	NamespaceName        string         `gorm:"column:NamespaceName" json:"namespace_name" form:"namespace_name"`
	BranchName           string         `gorm:"column:BranchName" json:"branch_name" form:"branch_name"`                        //暂时为发布类型的汉字
	ReleaseId            uint64         `gorm:"column:ReleaseId" json:"release_id" form:"release_id"`                           //保留字段
	PreviousReleaseId    uint64         `gorm:"column:PreviousReleaseId" json:"previous_release_id" form:"previous_release_id"` //保留字段
	Operation            uint64         `gorm:"column:Operation" json:"operation" form:"operation"`                             //发布类型
	OperationContext     []*models.Item `gorm:"column:OperationContext" json:"operation_context" form:"operation_context"`
	ReleaseContext       []*models.Item `gorm:"column:ReleaseContext" json:"release_context" form:"release_context"`
	DataChange_CreatedBy string         `gorm:"column:DataChange_CreatedBy" json:"data_change_created_by" form:"data_change_created_by"`
	//DataChange_LastModifiedBy string    `gorm:"column:DataChange_LastModifiedBy" json:"data_change_last_modified_by" form:"data_change_last_modified_by"`
	DataChange_CreatedTime time.Time `gorm:"column:DataChange_CreatedTime" json:"data_change_created_time" form:"data_change_created_time"`
	//DataChange_LastTime       time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}
