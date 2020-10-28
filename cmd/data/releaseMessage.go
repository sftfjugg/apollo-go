package main

//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
//	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
//	"go.didapinche.com/time"
//)
//
//func main() {
//
//	db1, err := gorm.Open("mysql", dsn1)
//	if err != nil {
//		panic(err)
//	}
//	db2, err := gorm.Open("mysql", dsn2)
//	if err != nil {
//		panic(err)
//	}
//
//	releaseMessage := make([]*models.ReleaseMessage, 0)
//	db1.Find(&releaseMessage)
//	//ReleaseMessage数据导入
//	for i, _ := range releaseMessage {
//		releaseMessage[i].Id = 0
//		releaseMessage[i].DataChange_LastTime = time.Now()
//		db := db2.Begin()
//		if err := db.Create(releaseMessage[i]).Error; err != nil {
//			db.Rollback()
//			fmt.Println(releaseMessage[i])
//		}
//		db.Commit()
//	}
//	fmt.Println("ReleaseMessage successed")
//
//}
