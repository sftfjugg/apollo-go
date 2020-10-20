package models

const FavoriteTableName = "Favorite"

//应用收藏表2
type Favorite struct {
	Base     Base   `gorm:"embedded"`
	UserId   string `gorm:"column:UserId" json:"user_id" form:"user_id"`
	AppId    string `gorm:"column:AppId" json:"app_id" form:"app_id"`
	Position uint64 `gorm:"column:Position" json:"position" form:"position"`
}

func (Favorite) TableName() string {
	return FavoriteTableName
}
