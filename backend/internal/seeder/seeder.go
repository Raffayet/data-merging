package seeder

import (
	"context"
	"log"
	"time"

	"github.com/Raffayet/data-merging/backend/internal/repository"
)

func GenerateDemoData(profileRepo *repository.MongoProfileRepository) {
	// Create demo data
	CleanDemoData(profileRepo)
	GenerateProfiles(profileRepo)
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
