package controller

import (
	"golang.org/x/net/websocket"

	log "github.com/inconshreveable/log15"
)

//WebsocketCtrl websocket  message
func WebsocketCtrl(ws *websocket.Conn) {
	for {
		log.Info("=== Start ===")
		msg := make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Error("Receive", "error", err)
		}
		log.Info("Receive ", "Message", string(msg[:n]))

		sendMsg := "[Roger that]"
		m, err := ws.Write([]byte(sendMsg))
		if err != nil {
			log.Error("Send", "error", err)
		}
		log.Info("Send ", "Message", string(sendMsg[:m]))
	}
}
