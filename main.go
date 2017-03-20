package main

import (
	"html/template"

	ctrl "github.com/cdyue/websocket/controller"
	"github.com/cdyue/websocket/handler"
	"golang.org/x/net/websocket"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	templ := template.Must(template.New("projectViews").ParseGlob("templates/*.tmpl"))
	r.SetHTMLTemplate(templ)

	r.Use(handler.Log()) //log15日志
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.POST("/user", gin.Bind(ctrl.User{}), ctrl.CreateUser)

	r.GET("/websocket", func(c *gin.Context) {
		handler := websocket.Handler(ctrl.WebsocketCtrl)
		handler.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "ws.tmpl", nil)
	})
	r.Run(":4000")
}
