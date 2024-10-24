package domain

type Address struct {
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	AddressLine string `json:"address_line" bson:"address_line"`
	PostalCode  string `json:"postal_code" bson:"postal_code"`
}
