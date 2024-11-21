package main

import (
	"1008001/splitwiser/internal/store"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

func main() {
	// path := filepath.Join(os.TempDir(), "db.sqlite")
	path := "db.sqlite"
	db, err := store.Init(path)
	if err != nil {
		trace := string(debug.Stack())
		slog.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
	defer db.Close()

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: routes(),
	}

	err = server.ListenAndServe()
	slog.Error(err.Error())
}
