package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

func HttpExternal(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	span, ctx := apm.StartSpan(c.Request.Context(), "go-app-external-request", "custom")
	defer span.End()
	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/info", os.Getenv("GO_INFO_URL")),
		nil,
	)

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

	c.JSON(200, gin.H{
		"message": "OK",
		"body": string(body),
		"http_status": http.StatusOK,
	})
}

