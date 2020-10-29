package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

func Release(db1 *gorm.DB, db2 *gorm.DB) {
	release := make([]*models.Release, 0)
	db1.Raw("select Id,ReleaseKey,Comment,AppId,ClusterName,NamespaceName,Configurations,DataChange_CreatedBy,DataChange_LastModifiedBy from `Release` where IsDeleted=0 and IsAbandoned=0 and Id in (select max(Id) from `Release`  group by AppId,NamespaceName) ORDER BY Id;").Scan(&release)
	for i, _ := range release {
		release[i].Id = 0
		release[i].DataChange_CreatedTime = time.Now()
		release[i].DataChange_LastTime = time.Now()
		release[i].IsDeleted = false
		release[i].IsAbandoned = false
		db := db2.Begin()
		if err := db.Create(release[i]).Error; err != nil {
			db.Rollback()
			fmt.Println("release导入失败，失败原因无法insert")
			fmt.Print(release[i])
		}
		db.Commit()
		fmt.Println("release导入成功")
		fmt.Print(release[i])
	}
	fmt.Println("AppNamespaceId successed")

}
