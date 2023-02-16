package middleware

import (
	"errors"
	"fmt"
	"strings"

	commons "fast-food-api-client/commons"
	"fast-food-api-client/configs"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	helpers "fast-food-api-client/helpers"
	resources "fast-food-api-client/resources"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var logAuth = coreLogger.Logger(constants.AppName, constants.StructAuthMiddleware)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logAuth.Info("[BEGIN] authMiddleware...")

		var res *commons.ResponseModel

		req := c.Request
		// get token from header
		bearerToken := req.Header.Get("Authorization")
		logAuth.Info("Token = ", "["+bearerToken+"]")

		var tokenString string
		if !helpers.IsNil(bearerToken) {
			tokenString = bearerToken
		}
		if strings.Contains(bearerToken, "Bearer") {
			tokenString = strings.Split(bearerToken, " ")[1]
		}
		// Check exists token in headers
		if helpers.IsNil(tokenString) {
			logAuth.Error("Token Not Found")
			err := errors.New("Not Found Token In Header With Request Path: " + req.URL.String())
			res = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestTokenNotFoundError)
			res.ResponseCommon(c)
			c.Abort()
			return
		}
		// verify token
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logAuth.Error("Token Is Invalid", token.Header["alg"])
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(configs.AppSecretKey), nil
		})
		// handler token error
		if err != nil {
			logAuth.Error("[END] authMiddleware... With Error Parse = ", err)
			res = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestTokenIsInValidError)
			res.ResponseCommon(c)
			c.Abort()
			return
		} else {
			logAuth.Info("[END] authMiddleware... Without Error")
			c.Next()
		}
	}
}
