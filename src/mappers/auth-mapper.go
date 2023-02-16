package mappers

import (
	"fast-food-api-client/helpers"
	"fast-food-api-client/src/dtos"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMapper struct{}

func (mapper *AuthMapper) ToAuthRegisterDTO(m *mongo.InsertOneResult) *dtos.AuthRegisterDTO {
	if helpers.IsEmpty(m) {
		return nil
	}
	return &dtos.AuthRegisterDTO{
		ID: m.InsertedID,
	}
}

func (mapper *AuthMapper) ToAuthLoginDTO(token string) *dtos.AuthLoginDTO {
	return &dtos.AuthLoginDTO{
		Token: token,
	}
}
