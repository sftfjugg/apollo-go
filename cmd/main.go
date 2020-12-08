package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.didapinche.com/foundation/agollo"
)

func test() (errs error) {
	defer func() {
		if err := recover(); err != nil {
			errs = errors.New("test")
		}
	}()

	a := 1
	a++
	fmt.Println("start")
	panic("Big Error")
	fmt.Println("stop")
	return nil
}
func main() {
	i := 0
	i += 1
	i = +2

	//idc := os.Getenv("IDC")
	//if idc != "" {
	//	viper.Set("apollo.cluster", idc)
	//}
	//appId := os.Getenv("APP_ID")
	//if appId != "" {
	//	viper.Set("apollo.appId", appId)
	//}
	//ip := os.Getenv("APOLLO_META")
	//if appId != "" {
	//	viper.Set("apollo.ip", ip)
	//}
	//namespace := os.Getenv("APOLLO_BOOTSTRAP_NAMESPACES")
	//if namespace != "" {
	//	viper.Set("apollo.namespaceName", namespace)
	//}
	//configPath := os.Getenv("APOLLO_CACHEDIR")
	//if namespace != "" {
	//	viper.Set("apollo.backupConfigPath", configPath)
	//}

	//viper.Set("apollo.appId", "apollo-test")
	//viper.Set("apollo.meta", "http://10.10.30.74:9090")
	//viper.Set("apollo.cluster", "default")
	//viper.Set("apollo.namespaceName", "application,test,bigdata.test")
	viper.SetConfigName("configs/app")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	agollo.Start()

	fmt.Println(viper.Get("server.port"))
	settings := viper.AllSettings()
	fmt.Println(settings)
	go func() {
		for {
			event := agollo.ListenChangeEvent()
			changeEvent := <-event
			bytes, _ := json.Marshal(changeEvent)
			fmt.Println("event:", string(bytes))
			fmt.Println(viper.AllSettings())
		}
	}()
	//p = viper.Get("db.password")
	//fmt.Println(p)

	config := agollo.GetCurrentApolloConfig()
	fmt.Println(config)
	select {}
	//cache := agollo.GetApolloConfigCache()
	//it := cache.NewIterator()
	//for i := 0; i < int(cache.EntryCount()); i++ {
	//	entry := it.Next()
	//	if entry == nil {
	//		continue
	//	}
	//	fmt.Printf("key : %s , value : %s ", string(entry.Key), string(entry.Value))
	//}
}
