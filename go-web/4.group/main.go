package main

import (
	"lege"
	"net/http"
)

func main() {
	r := lege.New()
	r.GET("/index", func(c *lege.Context) {
		c.HTML(http.StatusOK, "<h1>Index page</h1>")
	})
	v1 := r.Group("/v1")

	{
		v1.GET("/", func(c *lege.Context) {
			c.HTML(http.StatusOK, "<h1>hello lege<h1>")
		})

		v1.GET("/hello", func(c *lege.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *lege.Context) {
			// expect /hello/legektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *lege.Context) {
			c.JSON(http.StatusOK, lege.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	r.Run(":9999")
}
