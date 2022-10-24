package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	// 开启goroutine 把0-100写入到ch1通道中
	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		close(c1)
	}()

	// 开启goroutine 从ch1中取值，值的平方赋值给 ch2
	go func() {
		for {
			i, ok := <-c1
			if ok {
				c2 <- i * i
			} else {
				break
			}
		}

		close(c2)
	}()

	for i := range c2 {
		fmt.Println(i)
	}
}
