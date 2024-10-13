package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Profile represents a user profile in the system
type Profile struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
	Age   int                `json:"age" bson:"age"`
}
