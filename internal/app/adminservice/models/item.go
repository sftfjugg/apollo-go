package models

type Item struct {
	Key   string `gorm:"column:Key" json:"key" form:"value"`
	Value string `gorm:"column:Value" json:"value" form:"value"`
}
