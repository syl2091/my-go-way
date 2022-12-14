package model

import "errors"

type Arith struct {
}

// 请求结构体
type ArithRequest struct {
	A int
	B int
}

// 响应结构体
type ArithResponse struct {
	Pro int //乘积
	Quo int //商
	Rem int //余数
}

// 乘积方法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}
