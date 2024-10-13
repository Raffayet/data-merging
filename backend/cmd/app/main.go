package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Raffayet/data-merging/backend/internal/api"
	"github.com/Raffayet/data-merging/backend/internal/domain"
	"github.com/Raffayet/data-merging/backend/internal/repository"
	"github.com/Raffayet/data-merging/backend/internal/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	client := setupMongoDb()

	// Defer disconnecting the client until the application is shutting down
	defer client.Disconnect(context.Background())

	// Create repository, service, and handler
	profileRepo := repository.NewMongoProfileRepository(client)
	profileService := service.NewProfileService(profileRepo)
	profileHandler := api.NewProfileHandler(profileService)

	demoDataSeeder(profileRepo)
	router := setupRouter(profileHandler)

	startServer(router)
}

func startServer(router *mux.Router) {
	cors := setupCorsPolicy()

	// Get API URI from environment variable
	apiAddress := os.Getenv("API_URI_SHORTENED")
	if apiAddress == "" {
		log.Println("API_URI not set, using default :8000")
		apiAddress = ":8000"
	}

	// Start the HTTP server with CORS middleware
	log.Printf("Server is running on %s", apiAddress)
	err := http.ListenAndServe(apiAddress, cors(router))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func setupRouter(profileHandler *api.ProfileHandler) *mux.Router {
	// Setup HTTP routing
	router := mux.NewRouter()
	router.HandleFunc("/profiles", profileHandler.GetProfilesHandler).Methods("GET")
	return router
}

func setupMongoDb() *mongo.Client {
	// Get MongoDB URI from .env
	mongoURI := os.Getenv("MONGO_URI")
	client, err := repository.InitializeMongoClient(mongoURI)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	return client
}

func setupCorsPolicy() (corsMiddleware func(http.Handler) http.Handler) {
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

func demoDataSeeder(profileRepo *repository.MongoProfileRepository) {
	// Create demo data
	CleanDemoData(profileRepo)
	GenerateDemoData(profileRepo)
}

// CleanDemoData removes the demo profiles from MongoDB
func CleanDemoData(repo *repository.MongoProfileRepository) {
	// Use the MongoDB client to drop the entire database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.Client().Database("data_merging").Drop(ctx)
	if err != nil {
		log.Println("Error cleaning the database:", err)
	} else {
		log.Println("Database cleaned successfully, all collections dropped.")
	}
}

// GenerateDemoData creates and saves demo profiles to MongoDB
func GenerateDemoData(repo *repository.MongoProfileRepository) {
	profiles := []domain.Profile{
		{Name: "John Doe", Email: "john@example.com", Age: 30},
		{Name: "Jane Smith", Email: "jane@example.com", Age: 28},
		{Name: "Nikola S", Email: "nikola@example.com", Age: 23},
	}

	// Use repository to interact with MongoDB
	collection := repo.Client().Database("data_merging").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, profile := range profiles {
		_, err := collection.InsertOne(ctx, profile)
		if err != nil {
			log.Println("Error inserting demo profile:", err)
		}
	}
}
