package main

/*
(1) global middleware Logger
$ curl http://localhost:9999/
<h1>Hello lege</h1>

>>> log
2019/08/17 01:37:38 [200] / in 3.14µs
*/

/*
(2) global + group middleware
$ curl http://localhost:9999/v2/hello/legektutu
{"message":"Internal Server Error"}

>>> log
2019/08/17 01:38:48 [200] /v2/hello/legektutu in 61.467µs for group v2
2019/08/17 01:38:48 [200] /v2/hello/legektutu in 281µs
*/

import (
	lege "lege"
	"log"
	"net/http"
	"time"
)

func onlyForV2() lege.HandlerFunc {
	return func(c *lege.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := lege.New()
	r.Use(lege.Logger()) // global midlleware
	r.GET("/", func(c *lege.Context) {
		c.HTML(http.StatusOK, "<h1>Hello lege</h1>")
	})

	v2 := r.Group("/v2")
	//v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *lege.Context) {
			// expect /hello/legektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9998")
}
