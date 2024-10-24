package domain

type Organization struct {
	Base
	Name    string  `json:"name" bson:"name"`
	Address Address `json:"address" bson:"address"`
}
