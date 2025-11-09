package main

import (
	"fmt"
	"log"
	"net/http"
)

type RespHandler struct{}

func (h *RespHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("method: %s\n", r.Method)
	log.Printf("url:    %s\n", r.URL)
	fmt.Fprint(w, "Hello World")
}

func main() {
	handler := RespHandler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	log.Println("Server Listen at :8080 port...")
	server.ListenAndServe()
}
