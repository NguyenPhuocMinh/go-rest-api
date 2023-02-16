package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthLoginModel struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type AuthComparePasswordModel struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type AuthRegisterModel struct {
	FirstName string    `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string    `json:"lastName" bson:"lastName" validate:"required"`
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	Password  string    `json:"password" bson:"password" validate:"required"`
	Slug      string    `json:"slug,omitempty" bson:"slug"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}
