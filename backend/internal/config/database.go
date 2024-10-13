package config

import (
	"log"
	"os"

	"github.com/Raffayet/data-merging/backend/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupMongoDb() *mongo.Client {
	// Get MongoDB URI from .env
	mongoURI := os.Getenv("MONGO_URI")
	client, err := repository.InitializeMongoClient(mongoURI)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	return client
}
