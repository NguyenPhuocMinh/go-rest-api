package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerDTO struct {
	ID        primitive.ObjectID `json:"id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
}

type CustomerCreateDTO struct {
	ID interface{} `json:"id"`
}
