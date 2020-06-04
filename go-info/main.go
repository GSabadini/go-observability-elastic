package main

import (
	"database/sql"
	"expvar"
	"log"
	"net/http"

	"github.com/GSabadini/go-apm-elastic/go-info/handler"
	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

var counter *expvar.Int

func init() {
	counter = expvar.NewInt("counter")
}

func main() {
	//Connect database
	db, err := newPostgresHandler()
	if err != nil {
		log.Println(err)
	}

	//Gorilla router
	router := mux.NewRouter()

	//APM Agent
	apmgorilla.Instrument(router)

	//Route used by APM for query metrics
	router.Handle("/query/{name}", handler.Query(db))

	//Route used by APM for inbound http request metrics
	router.HandleFunc("/info", handler.Info)

	//Route used by APM for external http request metrics
	router.HandleFunc("/http-external", handler.HttpExternal)

	//Route used by Heartbeat for uptime metrics
	router.HandleFunc("/health", handler.Health)

	//Route used by Metricbeat for golang metrics
	router.Handle("/debug/vars", http.DefaultServeMux)

	log.Println("Start HTTP server :3001")
	if err := http.ListenAndServe(":3001", router); err != nil {
		panic(err)
	}
}

func newPostgresHandler() (*sql.DB, error) {
	dataSource := "host=postgres port=5432 user=dev dbname=info sslmode=disable password=dev"
	db, err := apmsql.Open("postgres", dataSource)
	if err != nil {
		return &sql.DB{}, err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS stats (name TEXT PRIMARY KEY, count INTEGER);"); err != nil {
		return &sql.DB{}, err
	}

	log.Println("Database connected")

	return db, nil
}