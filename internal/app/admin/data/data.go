package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.uber.org/zap"
)

//data包只是在老版本与新版本apollo数据迁移过程中使用，不需要进行代码检查
var dsn1 = "dida_plat:4yh4BhbPfwm@tcp(192.168.200.1:3026)/dida_apollo_config?charset=utf8mb4&parseTime=True&loc=Local"

//需要导入的mysql
var dsn2 = "dida_plat:4yh4BhbPfwm@tcp(192.168.200.1:3026)/dida_apollo_plus_config?charset=utf8mb4&parseTime=True&loc=Local"

//
var log, _ = zap.NewDevelopment()

//老版本apollo数据导入
func ImportData(dbs1, dbs2 string) {

	if dbs1 != "" {
		dsn1 = dbs1
	}

	if dbs2 != "" {
		dsn2 = dbs2
	}

	db1, err := gorm.Open("mysql", dsn1)
	if err != nil {
		panic(err)
	}
	db2, err := gorm.Open("mysql", dsn2)
	if err != nil {
		panic(err)
	}
	AppNamespace(db1, db2)
	Item(db1, db2)
	Release(db1, db2)
	ReleaseMessage(db1, db2)
	appNamespace := make([]*models.AppNamespace, 0)
	if err := db1.Raw("select AppId,Name from AppNamespace where IsPublic=1 and IsDeleted=0").Scan(&appNamespace).Error; err != nil {
		log.Error("查询公共配置失败")
	}
	for i := range appNamespace {
		if err := db2.Exec("update AppNamespace set AppId='public_global_config',IsPublic=1 where  AppId=? and Name=?;", appNamespace[i].AppId, appNamespace[i].Name).Error; err != nil {
			log.Error("修改公共配置失败:" + fmt.Sprint(appNamespace[i]))
		}
		log.Info("修改公共配置成功:" + fmt.Sprint(appNamespace[i]))
	}
	for i := range appNamespace {
		if err := db2.Exec("update `Release` set AppId='public_global_config' where  AppId=? and NamespaceName=?;", appNamespace[i].AppId, appNamespace[i].Name).Error; err != nil {
			log.Error("修改公共配置失败:" + fmt.Sprint(appNamespace[i]))
		}
		log.Info("修改公共配置成功:" + fmt.Sprint(appNamespace[i]))
	}
	log.Info("所有配置修改成功")

}

func UpadteDate(dbs2 string) {
	if dbs2 != "" {
		dsn2 = dbs2
	}
	db2, err := gorm.Open("mysql", dsn2)
	if err != nil {
		panic(err)
	}

	if err := db2.Exec("update `Release` set LaneName='default'").Error; err != nil {
		log.Error("修改公共配置失败:")
	}
	log.Info("修改公共配置成功:")
	if err := db2.Exec("Delete from `AppNamespace` where ClusterName!='default'").Error; err != nil {
		log.Error("修改公共配置失败:")
	}
	log.Info("修改公共配置成功:")
	if err := db2.Exec("update `AppNamespace` set LaneName='default'").Error; err != nil {
		log.Error("修改公共配置失败:")
	}
	log.Info("修改公共配置成功:")

}
