package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		w.Write([]byte("Hello world!"))
	})
	http.ListenAndServe(":3000", mux)
}

// $ go run HTTP/httpServer.go
// Request received

// $ curl -i localhost:3000/
// HTTP/1.1 200 OK
// Date: Wed, 22 Mar 2023 03:18:26 GMT
// Content-Length: 12
// Content-Type: text/plain; charset=utf-8

// Hello world!
