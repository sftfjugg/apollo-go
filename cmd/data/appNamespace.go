package main

//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
//	"go.didapinche.com/time"
//)
//
//func main() {
//	db1, err := gorm.Open("mysql", dsn1)
//	if err!=nil{
//		panic(err)
//	}
//	db2, err := gorm.Open("mysql", dsn2)
//	if err!=nil{
//		panic(err)
//	}
//
//	appNamespace:=make([]*models.AppNamespace,0)
//	db1.Raw("select AppId,ClusterName,NamespaceName Name,DataChange_CreatedBy,DataChange_LastModifiedBy from Namespace where IsDeleted=0;").Scan(&appNamespace)
//	//AppNamespace数据导入
//	for i,_:=range appNamespace{
//	db:=db2.Begin()
//		appNamespace[i].Id=0
//		appNamespace[i].DataChange_CreatedTime=time.Now()
//		appNamespace[i].DataChange_LastTime=time.Now()
//		appNamespace[i].IsDeleted=false
//		if appNamespace[i].ClusterName=="default"{
//			appNamespace[i].IsPublic=true
//			appNamespace[i].LaneName="主版本"
//		} else {
//			appNamespace[i].IsPublic=false
//			appNamespace[i].LaneName=appNamespace[i].ClusterName
//		}
//		if appNamespace[i].Name=="application"{
//			appNamespace[i].Format="服务"
//		}else {
//			appNamespace[i].Format="业务"
//		}
//		if err:=db.Create(appNamespace[i]).Error;err!=nil{
//			fmt.Println(appNamespace[i])
//			db.Rollback()
//		}
//	db.Commit()
//	}
//	fmt.Println("ReleaseMessage successed")
//}
