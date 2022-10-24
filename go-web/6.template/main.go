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
	"fmt"
	"html/template"
	"lege"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := lege.New()
	r.Use(lege.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "legektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *lege.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *lege.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", lege.H{
			"title":  "lege",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *lege.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", lege.H{
			"title": "lege",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
