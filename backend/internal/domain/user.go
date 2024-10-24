package domain

type User struct {
	Base
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"first_name" bson:"last_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}
