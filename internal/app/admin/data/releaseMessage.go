package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

func ReleaseMessage(db1 *gorm.DB, db2 *gorm.DB) {
	releaseMessage := make([]*models.ReleaseMessage, 0)
	db1.Find(&releaseMessage)
	//ReleaseMessage数据导入
	for i, _ := range releaseMessage {
		releaseMessage[i].Id = 0
		releaseMessage[i].DataChange_LastTime = time.Now()
		db := db2.Begin()
		if err := db.Create(releaseMessage[i]).Error; err != nil {
			db.Rollback()
			fmt.Println("releaseMessage倒入失败，失败原因无法insert")
			fmt.Println(releaseMessage[i])
		}
		db.Commit()
		fmt.Println("releaseMessage倒入成功")
		fmt.Println(releaseMessage[i])
	}
	fmt.Println("ReleaseMessage successed")

}
