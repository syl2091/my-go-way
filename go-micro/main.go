package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	proto "micro/helloworld/proto"
)

type Greeter struct {
}

// 实现 .pb.micro.go中的Hello方法 定义rsp的返回值
func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	//定义服务
	service := micro.NewService(micro.Name("greeter"))
	//服务初始化
	service.Init()
	//注册handler
	err := proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		return
	}
	//启动服务

	if err = service.Run(); err != nil {
		fmt.Println(err)
	}
}
