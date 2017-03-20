package controller

import (
	"golang.org/x/net/websocket"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

var users map[string]*websocket.Conn

//WebsocketCtrl websocket  message
func WebsocketCtrl(ws *websocket.Conn) {
	for {
		log.Info("=== Start ===")
		msg := make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Error("Receive", "error", err)
		}
		log.Info("Receive ", "Message", string(msg[:n]), "conn", ws)

		sendMsg := "[Roger that]"
		m, err := ws.Write([]byte(sendMsg))
		if err != nil {
			log.Error("Send", "error", err)
		}
		log.Info("Send ", "Message", string(sendMsg[:m]))
	}
}

//User User info
type User struct {
	Name string `json:"name" form:"name"`
}

//CreateUser createuser
func CreateUser(c *gin.Context) {
	args := c.MustGet(gin.BindKey).(*User)
	if args.Name == "" {
		c.JSON(404, struct{}{})
		return
	}
}
