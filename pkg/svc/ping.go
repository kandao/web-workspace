package svc

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	slog.Info("PingPong!!")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
