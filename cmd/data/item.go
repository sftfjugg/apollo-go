package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

func main() {
	db1, err := gorm.Open("mysql", dsn1)
	if err != nil {
		panic(err)
	}
	db2, err := gorm.Open("mysql", dsn2)
	if err != nil {
		panic(err)
	}
	item := make([]*models.Item, 0)
	db1.Raw("select NamespaceId,`Key`,Value,Comment `Describe`,DataChange_CreatedBy,DataChange_LastModifiedBy from `Item` where IsDeleted=0").Scan(&item)
	for i, _ := range item {
		item[i].Id = 0
		item[i].DataChange_LastTime = time.Now()
		item[i].DataChange_CreatedTime = time.Now()
		item[i].ReleaseValue = item[i].Value
		item[i].Status = 1
		item[i].Comment = "老系统迁移"
		item[i].IsDeleted = false
		appNamespace := new(models.AppNamespace)
		if err := db1.Raw("select AppId,ClusterName,NamespaceName Name from Namespace where IsDeleted=0 and Id=?;", item[i].NamespaceId).Scan(&appNamespace); err != nil {
			fmt.Println(item[i])
		}
		if err := db2.Raw("select Id from AppNamespace where IsDeleted=0 and AppId=? and ClusterName=? and Name=?;", appNamespace.AppId, appNamespace.ClusterName, appNamespace.Name).Scan(&appNamespace); err != nil {
			fmt.Println(item[i])
		}
		item[i].NamespaceId = appNamespace.Id
		db := db2.Begin()
		if err := db.Create(item[i]).Error; err != nil {
			db.Rollback()
			fmt.Println(item[i])
		}
		db.Commit()
	}
	fmt.Println("AppNamespaceId successed")

}
