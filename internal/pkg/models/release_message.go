package models

import "go.didapinche.com/time"

const ReleaseMessageTableName = "ReleaseMessage"

//发布消息 1,这个表是服务端实现推模式的关键，通过该表建立一个消息中间件，
//不停的拉去ID和Message，在Message相同的情况下判断最大ID，以此判断是否有最新值，客户端也会通过Id和Message来对比是否有最新值
type ReleaseMessage struct {
	Id                  uint64    `gorm:"column:Id" json:"id" form:"id"`
	Message             string    `gorm:"column:Message" json:"message" form:"message"`
	DataChange_LastTime time.Time `gorm:"column:DataChange_LastTime" json:"data_change_last_time" form:"data_change_last_time"`
}

func (ReleaseMessage) TableName() string {
	return ReleaseMessageTableName
}
