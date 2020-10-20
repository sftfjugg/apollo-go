package models

const CommitTableName = "Commit"

//历史提交表 1
type Commit struct {
	Base          Base   `gorm:"embedded"`
	ChangeSets    string `gorm:"column:ChangeSets" json:"change_sets" form:"change_sets"`
	AppId         string `gorm:"column:AppId" json:"app_id" form:"app_id"`
	ClusterName   string `gorm:"column:ClusterName" json:"cluster_name" form:"cluster_name"`
	NamespaceName string `gorm:"column:NamespaceName" json:"namespace_name" form:"namespace_name"`
	Comment       string `gorm:"column:Comment" json:"comment" form:"comment"`
}

func (Commit) TableName() string {
	return CommitTableName
}
