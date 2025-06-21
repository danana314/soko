package safari

import (
	"1008001/soko/internal/safari/store"
	"net/http"
	"os"
)

// SetupSafariApp initializes and returns the safari app handler
func SetupSafariApp() (http.Handler, error) {
	// Get database path from environment variable
	dbPath := os.Getenv("SAFARI_DB_PATH")
	if dbPath == "" {
		dbPath = "data/safari.sqlite" // fallback default
	}
	
	// Initialize database
	_, err := store.Init(dbPath)
	if err != nil {
		return nil, err
	}
	
	// Note: We're not closing the DB here as it needs to stay open for the app lifetime
	// In a real production app, you'd want proper cleanup handling
	
	// Return the routes handler
	return routes(), nil
}