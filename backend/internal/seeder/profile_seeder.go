package seeder

import (
	"context"
	"log"
	"time"

	"github.com/Raffayet/data-merging/backend/internal/domain"
	"github.com/Raffayet/data-merging/backend/internal/repository"
)

// GenerateProfiles creates and saves demo profiles to MongoDB
func GenerateProfiles(repo *repository.MongoProfileRepository) {
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
