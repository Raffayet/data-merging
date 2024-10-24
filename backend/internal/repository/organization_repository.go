package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// ProfileRepository interface defines methods for profile data access
type OrganizationRepository interface {
	Client() *mongo.Client // Add a method to expose MongoDB client
}

type MongoOrganizationRepository struct {
	client *mongo.Client
}

func (repo *MongoOrganizationRepository) Client() *mongo.Client {
	return repo.client
}

func NewOrganizationRepository(client *mongo.Client) *MongoOrganizationRepository {
	return &MongoOrganizationRepository{client: client}
}
