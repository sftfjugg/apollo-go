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
			log.Error("修改公共配置失败:" + fmt.Sprint(releaseMessage[i]))
		}
		db.Commit()
		log.Error("修改公共配置成功:" + fmt.Sprint(releaseMessage[i]))
	}
	log.Info("ReleaseMessage successed")

}
