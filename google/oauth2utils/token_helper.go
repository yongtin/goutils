package oauth2utils

import (
	"context"
	"fmt"
	"net/http"
)

func getCodeFromLocalServer(ctx context.Context) (string, error) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})

	return "", fmt.Errorf("No code found")
}
