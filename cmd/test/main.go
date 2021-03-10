package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/uber/tchannel-go"
	"go.didapinche.com/goapi/apollo_thrift_service/v2"
	"go.didapinche.com/zeus-go/v2"
	"go.didapinche.com/zeus-go/v2/client"
	"strings"
	"time"
)

func main() {
	a := "  AP  P:123      4  1"
	as := strings.Split(a, " ")
	fmt.Println(as)
	var i int64
	i = time.Now().UnixNano() / 1e6
	fmt.Println(int(i))
	z, err := zeus.New("limos-app-name")
	if err != nil {
		panic(errors.Wrap(err, "failed to create zeus"))
	}
	// 2.构建客户端
	cli, err := client.New(z, "ApolloThriftService")
	if err != nil {
		panic(errors.Wrap(err, "failed to create HelloService client"))
	}
	apollo := apollo_thrift_service.NewTChanApolloThriftServiceClient(cli)
	//查询namespaceId
	appNamespce := new(apollo_thrift_service.AppNamespace)
	appNamespce.Name = "application"
	appNamespce.AppId = "apollo-test"
	appNamespce.Env = "1"
	appNamespce.Operator = "lihang"
	appNamespce.LaneName = "default"
	appNamespce.ClusterName = "default"
	ctx, _ := tchannel.NewContextBuilder(time.Second).Build()
	id, err := apollo.CreateOrFindAppNamespace(ctx, appNamespce)
	if err != nil {
		fmt.Println(err)
	}
	//修改配置
	item := new(apollo_thrift_service.Item)
	item.NamespaceId = id
	item.Key = "myName5"
	item.Value = "test"
	item.Env = "1"
	item.Operator = "lihang"
	ctx1, _ := tchannel.NewContextBuilder(time.Second).Build()
	if err := apollo.CreateOrUpdateItem(ctx1, item); err != nil {
		fmt.Println(err)
	}
	//发布
	release := new(apollo_thrift_service.Release)
	release.NamespaceId = id
	release.Env = "1"
	keys := make([]string, 0)
	keys = append(keys, item.Key)
	release.Keys = keys
	release.Operator = "lihang"
	ctx2, _ := tchannel.NewContextBuilder(time.Second).Build()
	if err := apollo.PublishNamespace(ctx2, release); err != nil {
		fmt.Println(err)
	}
}
