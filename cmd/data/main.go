package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dsn1 = "plat:mTAerlrufO@tcp(192.168.1.205:3600)/dida_apollo_config?charset=utf8mb4&parseTime=True&loc=Local"

var dsn2 = "plat:mTAerlrufO@tcp(192.168.1.205:3600)/plat_apollo_config?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db1, err := gorm.Open("mysql", dsn1)
	if err != nil {
		panic(err)
	}
	db2, err := gorm.Open("mysql", dsn2)
	if err != nil {
		panic(err)
	}
	AppNamespace(db1, db2)
	Item(db1, db2)
	Release(db1, db2)
	ReleaseMessage(db1, db2)

}
