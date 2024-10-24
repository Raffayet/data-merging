package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OrganizationId string             `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
	Archived       bool               `json:"archived" bson:"archived"`
	InsertDate     time.Time          `json:"insert_date,omitempty" bson:"insert_date,omitempty"`
	UpdateDate     time.Time          `json:"update_date,omitempty" bson:"update_date,omitempty"`
}
