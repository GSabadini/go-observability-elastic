package main

import (
	"database/sql"
	"log"

	"github.com/GSabadini/go-apm-elastic/go-app/handler"
	"github.com/gin-contrib/expvar"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmsql"
	//_ "go.elastic.co/apm/module/apmsql/sqlite3"
)

func main() {
	//Connect database
	//db, err := newSQLiteHandler()
	//if err != nil {
	//	log.Println(err)
	//}

	var (
		//Connect cache redis
		clientRedis = newClientRedis()

		//Start Gin
		engine = gin.New()
	)

	//APM Agent
	engine.Use(apmgin.Middleware(engine))

	//Route used by APM for inbound http request metrics
	engine.GET("/ping", handler.Ping)

	//Route used by APM for external http request metrics
	engine.GET("/http-external", handler.HttpExternal)

	//Route used by Heartbeat for uptime metrics
	engine.GET("/health", handler.Health)

	//Route used by APM for query metrics
	//engine.GET("/query/:name", handler.Query(db))

	//Route used by Metricbeat for redis metrics
	engine.GET("/cache/:key", handler.Cache(clientRedis))

	//Route used by Metricbeat for golang metrics
	engine.GET("/debug/vars", expvar.Handler())

	if err := engine.Run(":3000"); err != nil {
		panic(err)
	}
}

func newClientRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println("connected redis", pong, err)

	return client
}

func newSQLiteHandler() (*sql.DB, error) {
	db, err := apmsql.Open("sqlite3", "./test.db")
	if err != nil {
		return &sql.DB{}, err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS stats (name TEXT PRIMARY KEY, count INTEGER);"); err != nil {
		return &sql.DB{}, err
	}

	log.Println("Database connected")

	return db, nil
}

