package main

import (
	"fmt"
	"log"
	"net/rpc"
	model "rpc/model"
)

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("dialing error:", err)
	}

	req := model.ArithRequest{10, 20}
	var res model.ArithResponse

	err = conn.Call("Arith.Multiply", req, &res) //乘法运算
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	//除法运算
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d / %d = %d 余数是:%d", req.A, req.B, res.Quo, res.Rem)
}
