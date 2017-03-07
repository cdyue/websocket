package main

import (
	ctrl "github.com/cdyue/websocket/controller"
	"github.com/cdyue/websocket/handler"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(handler.Log()) //log15日志
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	rWs := r.Group("/ws")
	{
		rWs.POST("/test", ctrl.WebsocketCtrl)
	}

	r.Run(":4000")
}
