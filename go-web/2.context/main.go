package main

import (
	"lege"
	"net/http"
)

func main() {
	r := lege.New()
	r.GET("/", func(c *lege.Context) {
		c.HTML(http.StatusOK, "<h1>Hello lege</h1>")
	})
	r.GET("/hello", func(c *lege.Context) {
		// expect /hello?name=legektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *lege.Context) {
		c.JSON(http.StatusOK, lege.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
