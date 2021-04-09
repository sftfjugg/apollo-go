package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(2)
	fmt.Println(i)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			test()
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	go func(ticker *time.Ticker) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				test()
			}
		}()
		defer ticker.Stop()
		for range ticker.C {
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(2)
			fmt.Println(i)
		}
	}(ticker)
}

func main() {
	//list:=make(map[string]string,0)
	//list=nil
	//for i,_:=range  list{
	//	fmt.Print(i)
	//}

	test()
	time.Sleep(50 * time.Second)
	//test:=make(chan models.App,10)
	//go func() {
	//	var app models.App
	//	test <-app
	//}()
	//a:=<-test
	//fmt.Println(a)

}

//a := "  AP  P:123      4  1"
//as := strings.Split(a, " ")
//fmt.Println(as)
//var i int64
//i = time.Now().UnixNano() / 1e6
//fmt.Println(int(i))
//z, err := zeus.New("limos-app-name")
//if err != nil {
//panic(errors.Wrap(err, "failed to create zeus"))
//}
//// 2.构建客户端
//cli, err := client.New(z, "ApolloThriftService")
//if err != nil {
//panic(errors.Wrap(err, "failed to create HelloService client"))
//}
//apollo := apollo_thrift_service.NewTChanApolloThriftServiceClient(cli)
////查询namespaceId
//appNamespce := new(apollo_thrift_service.AppNamespace)
//appNamespce.Name = "application"
//appNamespce.AppId = "apollo-test"
//appNamespce.Env = "1"
//appNamespce.Operator = "lihang"
//appNamespce.LaneName = "default"
//appNamespce.ClusterName = "default"
//ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
//id, err := apollo.CreateOrFindAppNamespace(ctx, appNamespce)
//if err != nil {
//fmt.Println(err)
//}
////修改配置
//item := new(apollo_thrift_service.Item)
//item.NamespaceId = id
//item.Key = "myName5"
//item.Value = "test"
//item.Env = "1"
//item.Operator = "lihang"
//ctx1, _ := tchannel.NewContextBuilder(time.Second).Build()
//if err := apollo.CreateOrUpdateItem(ctx1, item); err != nil {
//fmt.Println(err)
//}
////发布
//release := new(apollo_thrift_service.Release)
//release.NamespaceId = id
//release.Env = "1"
//keys := make([]string, 0)
//keys = append(keys, item.Key)
//release.Keys = keys
//release.Operator = "lihang"
//ctx2, _ := tchannel.NewContextBuilder(time.Second).Build()
//if err := apollo.PublishNamespace(ctx2, release); err != nil {
//fmt.Println(err)
//}
