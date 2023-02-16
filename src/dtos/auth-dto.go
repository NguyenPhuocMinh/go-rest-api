package dtos

type AuthRegisterDTO struct {
	ID interface{} `json:"id"`
}

type AuthLoginDTO struct {
	Token string `json:"token"`
}
