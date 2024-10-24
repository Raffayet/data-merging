package domain

type Dataset struct {
	Base
	Content interface{} `json:"content" bson:"content"`
}
