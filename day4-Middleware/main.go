package main

import (
	"log"
	"net/http"
	"test04/gee"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		//start timer
		t := time.Now()
		//c.Fail(500,"Internal Server Error")
		log.Printf("[%d] %s in %v for group v2 ", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Lars!</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
