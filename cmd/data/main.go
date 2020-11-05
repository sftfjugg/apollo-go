package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

var dsn1 = "plat:mTAerlrufO@tcp(192.168.1.205:3600)/dida_apollo_config?charset=utf8mb4&parseTime=True&loc=Local"

var dsn2 = "plat:mTAerlrufO@tcp(192.168.1.205:3600)/dida_apollo_plus_config?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db1, err := gorm.Open("mysql", dsn1)
	if err != nil {
		panic(err)
	}
	db2, err := gorm.Open("mysql", dsn2)
	if err != nil {
		panic(err)
	}
	//AppNamespace(db1, db2)
	//Item(db1, db2)
	//Release(db1, db2)
	//ReleaseMessage(db1, db2)
	appNamespace := make([]*models.AppNamespace, 0)
	if err := db1.Raw("select AppId,Name from AppNamespace where IsPublic=1 and IsDeleted=0").Scan(&appNamespace).Error; err != nil {
		fmt.Println("查询公共配置失败")
	}
	for i := range appNamespace {
		if err := db2.Exec("update AppNamespace set AppId='public_global_config',IsPublic=1 where  AppId=? and Name=?;", appNamespace[i].AppId, appNamespace[i].Name).Error; err != nil {
			fmt.Println("修改公共配置失败:")
			fmt.Print(appNamespace[i])
		}
		fmt.Println("修改公共配置成功:")
		fmt.Print(appNamespace[i])
	}
	fmt.Println("所有数据完成")

}
