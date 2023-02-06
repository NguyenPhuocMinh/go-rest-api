package services

import (
	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	serviceV1 "fast-food-api-client/services/v1"

	"github.com/gin-gonic/gin"
)

type ResponseBaseService struct {
	MsgType     string
	MsgAction   string
	Service     FuncServiceHandler
	ServiceName string
}

type FuncServiceHandler func(c *gin.Context) *commons.ResponseModel

var logger = coreLogger.Logger(constants.AppName, constants.StructService)

func BaseService() []ResponseBaseService {
	logger.Info("BEGIN BaseService...")

	var services []ResponseBaseService

	// Auth Service
	authServices := []ResponseBaseService{
		{
			MsgType:     constants.MsgTypeAuth,
			MsgAction:   constants.MsgActionAuthLogin,
			Service:     serviceV1.AuthLogin,
			ServiceName: "AuthLogin",
		},
	}

	services = append(services, authServices...)

	logger.Info("END BaseService...")
	return services
}
