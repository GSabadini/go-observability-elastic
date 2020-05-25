package main

import (
	"context"
	//"database/sql"
	"fmt"

	"net/http"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmhttp"

	//"go.elastic.co/apm/module/apmsql"
	//_ "go.elastic.co/apm/module/apmsql/sqlite3"

	"github.com/gin-contrib/expvar"
	"github.com/gin-gonic/gin"
	//_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine := gin.New()

	engine.Use(apmgin.Middleware(engine))

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/http-client", func(c *gin.Context) {
		monitorRequest(c.Request.Context())
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	//engine.GET("/query", func(c *gin.Context) {
	//	var vars = c.Request.URL.Query()
	//	var name = vars.Get("name")
	//	requestCount, err := updateRequestCount(c.Request.Context(), name)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	c.JSON(200, gin.H{
	//		"message": "success",
	//		"count":   requestCount,
	//	})
	//})

	engine.GET("/debug/vars", expvar.Handler())

	engine.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func monitorRequest(ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "monitorRequest", "custom")
	defer span.End()
	req, _ := http.NewRequest(http.MethodGet, "https://enzzfdy8g188g.x.pipedream.net", nil)

	// Faça instrumentação do cliente HTTP e adicione o contexto circundante à solicitação.
	// Isso fará com que uma duração seja gerada para a solicitação HTTP de saída, incluindo
	// um cabeçalho de rastreamento distribuído para continuar o rastreamento no serviço de destino.
	client := apmhttp.WrapClient(http.DefaultClient)
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Println(err)
		apm.CaptureError(ctx, err).Send()
	}
	defer resp.Body.Close()
}

//var db *sql.DB

//func connectDB() {
//	var err error
//	db, err = apmsql.Open("sqlite3", ":memory:")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if _, err := db.Exec("CREATE TABLE stats (name TEXT PRIMARY KEY, count INTEGER);"); err != nil {
//		log.Fatal(err)
//	}
//
//}

// updateRequestCount incrementa uma contagem para o nome no banco de dados, retornando a nova contagem.
//func updateRequestCount(ctx context.Context, name string) (int, error) {
//	tx, err := db.BeginTx(ctx, nil)
//	if err != nil {
//		return -1, err
//	}
//	row := tx.QueryRowContext(ctx, "SELECT count FROM stats WHERE name=?", name)
//	var count int
//	switch err := row.Scan(&count); err {
//	case nil:
//		count++
//		if _, err := tx.ExecContext(ctx, "UPDATE stats SET count=? WHERE name=?", count, name); err != nil {
//			return -1, err
//		}
//	case sql.ErrNoRows:
//		count = 1
//		if _, err := tx.ExecContext(ctx, "INSERT INTO stats (name, count) VALUES (?, ?)", name, count); err != nil {
//			return -1, err
//		}
//	default:
//		return -1, err
//	}
//
//	return count, tx.Commit()
//}
