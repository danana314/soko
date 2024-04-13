package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handleIndex)

	http.ListenAndServe(":8080", router)
}
