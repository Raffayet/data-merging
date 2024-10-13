package repository

import (
	"context"
	"time"

	"github.com/Raffayet/data-merging/backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProfileRepository interface defines methods for profile data access
type ProfileRepository interface {
	FetchProfiles(ctx context.Context) ([]domain.Profile, error)
	Client() *mongo.Client // Add a method to expose MongoDB client
}

// MongoProfileRepository implements ProfileRepository with MongoDB
type MongoProfileRepository struct {
	client *mongo.Client
}

func NewMongoProfileRepository(client *mongo.Client) *MongoProfileRepository {
	return &MongoProfileRepository{client: client}
}

func (repo *MongoProfileRepository) Client() *mongo.Client {
	return repo.client
}

// FetchProfiles retrieves all profiles from MongoDB
func (repo *MongoProfileRepository) FetchProfiles(ctx context.Context) ([]domain.Profile, error) {
	var profiles []domain.Profile
	collection := repo.client.Database("data_merging").Collection("profiles")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var profile domain.Profile
		cursor.Decode(&profile)
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// InitializeMongoClient creates and returns a MongoDB client
func InitializeMongoClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set custom options, including max pool size
	clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(200) // Increase max pool size to 200

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure connection is established
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
