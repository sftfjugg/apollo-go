package models

const AppTableName = "App"

//项目 1 2
type App struct {
	Base       Base   `gorm:"embedded"`
	AppId      string `gorm:"column:AppId" json:"app_id" form:"app_id"`
	Name       string `gorm:"column:Name" json:"name" form:"name"`
	OrgId      string `gorm:"column:OrgId" json:"org_id" form:"org_id"` //部门id
	OrgName    string `gorm:"column:OrgName" json:"org_name" form:"org_name"`
	OwnerName  string `gorm:"column:OwnerName" json:"owner_name" form:"owner_name"`
	OwnerEmail string `gorm:"column:OwnerEmail" json:"owner_email" form:"owner_email"`
}

func (App) TableName() string {
	return AppTableName
}
