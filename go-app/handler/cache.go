package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.elastic.co/apm/module/apmgoredis"
)

func Cache(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := apmgoredis.Wrap(redisClient).WithContext(c.Request.Context())

		var key = c.Param("key")

		value, err := client.Get(key).Result()
		if err == redis.Nil {
			value = fmt.Sprintf("Key:%s does not exist", key)
			err := client.Set(key, fmt.Sprintf("Key:%s exist", key), 0).Err()
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}

		c.JSON(200, gin.H{
			"message":     "OK",
			"http_status": http.StatusOK,
			"key":       key,
			"result": value,
		})
	}
}