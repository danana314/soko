package main

import (
	"1008001/splitwiser/internal/store"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment-specific .env file
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	
	envFile := fmt.Sprintf(".env/%s", env)
	godotenv.Load(envFile)
	
	// Get database path from environment variable - panic if not set
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		panic("DB_PATH environment variable is required")
	}
	
	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback default
	}
	
	db, err := store.Init(dbPath)
	if err != nil {
		trace := string(debug.Stack())
		slog.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
	defer db.Close()

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", port),
		Handler: routes(),
	}

	err = server.ListenAndServe()
	slog.Error(err.Error())
}
