package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

func HttpExternal(w http.ResponseWriter, r *http.Request) {
	span, ctx := apm.StartSpan(r.Context(), "go-info-request-external", "custom")
	defer span.End()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ping", os.Getenv("GO_APP_URL")), nil)

	client := apmhttp.WrapClient(http.DefaultClient)
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
		apm.CaptureError(ctx, err).Send()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response = map[string]interface{}{
		"message": "OK",
		"body": string(body),
		"http_status": http.StatusOK,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

