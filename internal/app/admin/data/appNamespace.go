package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

func AppNamespace(db1 *gorm.DB, db2 *gorm.DB) {

	appNamespace := make([]*models.AppNamespace, 0)
	db1.Raw("select AppId,ClusterName,NamespaceName Name,DataChange_CreatedBy,DataChange_LastModifiedBy from Namespace where IsDeleted=0 and ClusterName='default';").Scan(&appNamespace)
	//AppNamespace数据导入
	for i, _ := range appNamespace {
		db := db2.Begin()
		appNamespace[i].Id = 0
		appNamespace[i].DataChange_CreatedTime = time.Now()
		appNamespace[i].DataChange_LastTime = time.Now()
		appNamespace[i].IsDeleted = false
		appNamespace[i].IsPublic = false
		if appNamespace[i].ClusterName == "default" {
			appNamespace[i].LaneName = "default"
		} else {
			appNamespace[i].LaneName = appNamespace[i].ClusterName
		}
		if appNamespace[i].Name == "application" {
			appNamespace[i].Format = "服务"
		} else {
			appNamespace[i].Format = "业务"
		}
		if err := db.Create(appNamespace[i]).Error; err != nil {
			log.Info("修改公共配置成功:" + fmt.Sprint(appNamespace[i]))
			db.Rollback()
		}
		log.Info("倒入成功" + fmt.Sprint(appNamespace[i]))
		db.Commit()
	}
	log.Info("ReleaseMessage successed")
}
