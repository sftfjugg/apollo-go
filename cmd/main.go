package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.didapinche.com/foundation/agollo"
	"os"
)

func main() {

	conn := os.Getenv("IDC")
	if conn != "" {
		viper.Set("apollo.cluster", "test")
	}
	//viper.Set("apollo.appId", "taxidetail-rs-service")
	//viper.Set("apollo.ip", "http://10.31.122.109:9090")
	//viper.Set("apollo.cluster", "test")
	//viper.Set("apollo.namespaceName", "test")
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
