package models

const UsersTableName = "Users"

//2
type Users struct {
	Id       uint64 `gorm:"column:Id" json:"id" form:"id"`
	Username string `gorm:"column:Username" json:"username" form:"username"`
	Password string `gorm:"column:Password" json:"password" form:"password"`
	Email    string `gorm:"column:Email" json:"email" form:"email"`
	Enabled  string `gorm:"column:Enabled" json:"enabled" form:"enabled"`
}

func (Users) TableName() string {
	return UsersTableName
}
