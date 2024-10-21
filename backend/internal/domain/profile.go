package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Profile represents a user profile in the system
type Profile struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string             `json:"first_name" bson:"first_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Address     string             `json:"address" bson:"address"`
}
