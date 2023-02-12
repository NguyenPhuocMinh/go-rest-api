package mappers

import (
	"fast-food-api-client/helpers"
	"fast-food-api-client/src/dtos"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerMapper struct{}

func (mapper *CustomerMapper) ToCreateDto(m *mongo.InsertOneResult) *dtos.CustomerCreateDTO {
	if helpers.IsEmpty(m) {
		return nil
	}
	return &dtos.CustomerCreateDTO{
		ID: m.InsertedID,
	}
}
