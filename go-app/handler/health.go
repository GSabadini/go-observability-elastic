package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(200, gin.H{
		"message":     "OK",
		"http_status": http.StatusOK,
	})
}
