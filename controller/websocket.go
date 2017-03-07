package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

//WebsocketCtrl websocket controller
func WebsocketCtrl(c *gin.Context) {

	log := c.MustGet("log").(log15.Logger)

	log.Info("=== Start ===")

	c.JSON(http.StatusOK, struct{}{})
}
