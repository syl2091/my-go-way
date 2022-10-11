package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"rpc/model"
)

func main() {
	//注册rpc服务
	rpc.Register(new(model.Arith))
	//采用http协议作为rpc载体
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("fatal error:", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection\n")

	//接收客户端请求 并发处理 jsonrpc
	for {
		conn, err := lis.Accept() //接收客户端连接请求
		if err != nil {
			continue
		}

		//并发处理客户端请求
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

	//常规启动http服务
	//http.Serve(lis,nil)
}
