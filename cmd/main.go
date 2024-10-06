package main

import (
	"1008001/splitwiser/internal/store"
	"net/http"
)

func main() {
	store.Init()
	// db, err := store.New("db.sqlite")
	// if err != nil {
	// trace := string(debug.Stack())
	// slog.Error(err.Error(), "trace", trace)
	// os.Exit(1)
	// }
	// defer db.Close()

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: routes(),
	}

	server.ListenAndServe()
}
