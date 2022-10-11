package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"rpc/model"
)

func main() {
	//注册rpc服务
	rpc.Register(new(model.Arith))
	//采用http协议作为rpc载体
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("fatal error:", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection\n")

	//常规启动http服务
	http.Serve(lis, nil)
}
