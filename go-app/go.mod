module github.com/GSabadini/go-apm-elastic/go-app

go 1.14

require (
	github.com/gin-contrib/expvar v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.0.0-beta.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	go.elastic.co/apm v1.8.0
	go.elastic.co/apm/module/apmgin v1.8.0
	go.elastic.co/apm/module/apmgoredis v1.8.0
	go.elastic.co/apm/module/apmhttp v1.8.0
	go.elastic.co/apm/module/apmsql v1.8.0
)
