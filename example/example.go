package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/yasid", func(w http.ResponseWriter, req *http.Request) {
		_ = req.ParseForm()
		var param = req.Form.Encode()
		var res = fmt.Sprintf("yasidcontent: %s\n", param)
		fmt.Fprintf(w, res)
	})

	http.ListenAndServe(":8080", nil)
}
