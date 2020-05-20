package main

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	engine := gin.New()
	engine.Use(apmgin.Middleware(engine))
	//r := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
