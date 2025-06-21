package main

import (
	"1008001/soko/cmd/apps/safari"
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
	
	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback default
	}
	
	// Initialize main router with subdirectory routing
	mainRouter := setupMainRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port), // Remove localhost for containerization
		Handler: mainRouter,
	}

	slog.Info("Starting server", "port", port)
	err := server.ListenAndServe()
	if err != nil {
		trace := string(debug.Stack())
		slog.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func setupMainRouter() http.Handler {
	mainRouter := http.NewServeMux()
	
	// Register safari app at /safari/ path
	safariHandler := setupSafariApp()
	mainRouter.Handle("/safari/", http.StripPrefix("/safari", safariHandler))
	
	// Root redirect to safari for now
	mainRouter.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/safari/", http.StatusMovedPermanently)
	})
	
	return mainRouter
}

func setupSafariApp() http.Handler {
	handler, err := safari.SetupSafariApp()
	if err != nil {
		slog.Error("Failed to setup safari app", "error", err)
		// Return error handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Safari app initialization failed", http.StatusInternalServerError)
		})
	}
	return handler
}