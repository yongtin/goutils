package simplehttp

import (
	"fmt"
	"net/http"
	"time"
)

func createServer(port string) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
