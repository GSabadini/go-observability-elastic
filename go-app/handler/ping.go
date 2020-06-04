package handler

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
