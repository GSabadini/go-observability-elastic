package main

import (
	"github.com/gin-contrib/expvar"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	engine := gin.New()

	engine.Use(apmgin.Middleware(engine))

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/debug/vars", expvar.Handler())

	engine.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
