package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.didapinche.com/agollo/v2"
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

	viper.SetConfigName("configs/app")
	viper.AddConfigPath("./")
	viper.Set("apollo.ip", "http://apollo-meta.didapinche.com")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	agollo.Start()
	settings := viper.AllSettings()
	fmt.Println(settings)
	//监听配置变更
	go func() {
		for {
			event := agollo.ListenChangeEvent()
			changeEvent := <-event
			bytes, _ := json.Marshal(changeEvent)
			fmt.Println("event:", string(bytes))
			fmt.Println(viper.AllSettings())
		}
	}()
	config := agollo.GetCurrentApolloConfig()
	fmt.Println(config)
	select {}
}
