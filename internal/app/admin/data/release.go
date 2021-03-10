package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

func Release(db1 *gorm.DB, db2 *gorm.DB) {
	release := make([]*models.Release, 0)
	db1.Raw("select Id,ReleaseKey,Comment,AppId,ClusterName,NamespaceName,Configurations,DataChange_CreatedBy,DataChange_LastModifiedBy from `Release` where IsDeleted=0 and IsAbandoned=0 and Id in (select max(Id) from `Release`  group by AppId,NamespaceName,ClusterName) ORDER BY Id;").Scan(&release)
	for i, _ := range release {
		release[i].Id = 0
		release[i].DataChange_CreatedTime = time.Now()
		release[i].DataChange_LastTime = time.Now()
		release[i].IsDeleted = false
		release[i].IsAbandoned = false
		release[i].LaneName = "default"
		db := db2.Begin()
		if err := db.Create(release[i]).Error; err != nil {
			db.Rollback()
			log.Error("修改公共配置成功:" + fmt.Sprint(release[i]))
		}
		db.Commit()
		log.Info("修改公共配置成功:" + fmt.Sprint(release[i]))
	}
	log.Info("AppNamespaceId successed")

}
