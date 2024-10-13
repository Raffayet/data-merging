package config

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func SetupCorsPolicy() (corsMiddleware func(http.Handler) http.Handler) {
	// Set up CORS to allow requests from http://localhost:3000 (React app)
	frontendURI := os.Getenv("FRONTEND_URI")
	if frontendURI == "" {
		log.Println("FRONTEND_URI not set, using default http://localhost:3000")
		frontendURI = "http://localhost:3000"
	}

	corsMiddleware = handlers.CORS(
		handlers.AllowedOrigins([]string{frontendURI}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	return corsMiddleware
}
