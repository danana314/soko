package main

import (
	"1008001/splitwiser/middleware"
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", handleIndex)

	server := http.Server{
		Addr:    "8080",
		Handler: middleware.Logging(router),
	}
	server.ListenAndServe()
}
