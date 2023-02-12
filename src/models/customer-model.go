package models

import (
	"time"
)

type CustomerModel struct {
	FirstName string    `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string    `json:"lastName" bson:"lastName" validate:"required"`
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	Password  string    `json:"password" bson:"password" validate:"required"`
	Slug      string    `json:"slug,omitempty" bson:"slug"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}
