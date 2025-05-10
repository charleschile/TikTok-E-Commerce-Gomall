package main

import (
	"context"
	"fmt"

	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/conf"
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/kitex_gen/pbapi"
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	// 注意教程中是kitex下的consul包而不是hertz下的
	// "github.com/hertz-contrib/registry/consul"

	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	// 找服务的名字需要改成demo_proto
	// 发起 RPC 调用时，解析器从 Consul 获取服务地址
	// 如果 Consul 返回多个健康的服务实例，resolver 会使用负载均衡策略选择一个
	client, err := echoservice.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	// 通过consul得到服务之后，直接通过rpc调用Echo服务
	res, err := client.Echo(context.Background(), &pbapi.Request{Message: "hello"})
	if err != nil {
		klog.Fatal(err)
	}
	fmt.Printf("%v", res)

}
