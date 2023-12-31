package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

func Item(db1 *gorm.DB, db2 *gorm.DB) {
	item := make([]*models.Item, 0)
	db1.Raw("select Id,NamespaceId,`Key`,Value,Comment `Describe`,DataChange_CreatedBy,DataChange_LastModifiedBy from `Item` where IsDeleted=0").Scan(&item)
	for i, _ := range item {
		appNamespace := new(models.AppNamespace)
		db1.Raw("select AppId,ClusterName,NamespaceName Name from Namespace where IsDeleted=0 and Id=?;", item[i].NamespaceId).Scan(&appNamespace)
		if appNamespace.AppId == "" {
			log.Info("修改公共配置成功:" + fmt.Sprint(item[i]))
		}
		db2.Raw("select Id from AppNamespace where IsDeleted=0 and AppId=? and ClusterName=? and Name=? and LaneName=?;", appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name, "default").Scan(&appNamespace)
		if appNamespace.Id == 0 {
			log.Error("修改公共配置失败:" + fmt.Sprint(item[i]))
		}
		item[i].NamespaceId = appNamespace.Id
		item[i].Id = 0
		item[i].DataChange_LastTime = time.Now()
		item[i].DataChange_CreatedTime = time.Now()
		item[i].ReleaseValue = item[i].Value
		item[i].Status = 1
		item[i].Comment = "老系统迁移"
		item[i].IsDeleted = false
		db := db2.Begin()
		if err := db.Create(item[i]).Error; err != nil {
			db.Rollback()
			log.Error("修改公共配置失败:" + fmt.Sprint(item[i]))
		}
		db.Commit()
		log.Info("修改公共配置成功:" + fmt.Sprint(item[i]))
	}
	log.Info("AppNamespaceId end")

}
