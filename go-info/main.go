package main

import (
	"encoding/json"
	"expvar"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

var counter *expvar.Int

func init() {
	counter = expvar.NewInt("counter")
}

func main() {
	router := mux.NewRouter()

	//APM Agent
	apmgorilla.Instrument(router)

	//Route used by APM for inbound http request metrics
	router.HandleFunc("/info", info)

	//Route used by Heartbeat for uptime metrics
	router.HandleFunc("/health", health)

	//Route used by Metricbeat for golang metrics
	router.Handle("/debug/vars", http.DefaultServeMux)

	log.Println("Start HTTP server :3001")
	if err := http.ListenAndServe(":3001", router); err != nil {
		panic(err)
	}
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func info(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()

	data := struct {
		Hostname string      `json:"hostname,omitempty"`
		IP       string      `json:"ip,omitempty"`
		Headers  http.Header `json:"headers,omitempty"`
		URL      string      `json:"url,omitempty"`
		Host     string      `json:"host,omitempty"`
		Method   string      `json:"method,omitempty"`
	}{
		Hostname: hostname,
		IP:       getIP(req),
		Headers:  req.Header,
		URL:      req.URL.RequestURI(),
		Host:     req.Host,
		Method:   req.Method,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getIP(req *http.Request) string {
	forwarded := req.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}

	return req.RemoteAddr
}
