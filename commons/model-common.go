package commons

import "github.com/golang-jwt/jwt/v4"

type ResponseModel struct {
	Data              interface{} `json:"data"`
	Error             *ErrorModel `json:"error"`
	Endpoint          string      `json:"endpoint"`
	Method            string      `json:"method"`
	StatusCode        int         `json:"statusCode"`
	MessageCode       string      `json:"messageCode"`
	MessageName       string      `json:"messageName"`
	MessageTranslator string      `json:"messageTranslator"`
	Pagination        interface{} `json:"pagination"`
}

type ErrorModel struct {
	ErrorDetail    string `json:"errorDetail"`
	ErrorFile      string `json:"errorFile"`
	ErrorComponent string `json:"errorComponent"`
}

type PayloadModel struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
