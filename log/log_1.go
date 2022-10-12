package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// 日志存储在文件中
func myDiv(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return x / y, nil
}

func main() {
	var x float64 = 18.2
	var y float64 = 0
	res, exception := myDiv(x, y)
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // 关闭文件

	log.SetOutput(file)
	log.Print(exception)
	fmt.Println(res)
}
