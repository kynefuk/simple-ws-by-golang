package main

import (
	"flag"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {

	flag.Parse()
	hub := newHub()
	go hub.run()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.File("home.html")
	})
	r.Any("/ws", func(c *gin.Context) {
		serveWs(hub, c)
	})

	r.Run()
}
