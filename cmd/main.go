package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.didapinche.com/foundation/agollo"
)

func main() {
	//viper.AutomaticEnv()
	viper.Set("apollo.appId", viper.GetString("taxidetail-rs-service"))
	viper.Set("apollo.meta", "http://10.31.77.101:9090")
	viper.Set("apollo.cluster", viper.GetString("test"))
	viper.Set("apollo.namespaceName", "test")
	viper.SetConfigName("app")
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
	//u = viper.Get("db.username")
	//fmt.Println(u)
	//p = viper.Get("db.password")
	//fmt.Println(p)

	select {}
	//config := agollo.GetCurrentApolloConfig()
	//cache := agollo.GetApolloConfigCache()
	//fmt.Println(config)
	//it := cache.NewIterator()
	//for i := 0; i < int(cache.EntryCount()); i++ {
	//	entry := it.Next()
	//	if entry == nil {
	//		continue
	//	}
	//	fmt.Printf("key : %s , value : %s ", string(entry.Key), string(entry.Value))
	//}
}
