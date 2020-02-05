package simplehttp

import (
	"io"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	hello2Handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, 2world!\n")
	}

	var httpserver = createServer("8080")
	httpserver.Handle("/hello", helloHandler)
	httpserver.Handle("/hello2", hello2Handler)

	// assert.Nil(t, err)
}
