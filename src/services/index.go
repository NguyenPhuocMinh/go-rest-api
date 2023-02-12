package services

import (
	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	serviceV1 "fast-food-api-client/src/services/v1"

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
	logger.Info("[BEGIN] BaseService...")

	var services []ResponseBaseService

	commonServices := []ResponseBaseService{
		{
			MsgType:     constants.MsgTypeCommon,
			MsgAction:   constants.MsgActionTest,
			Service:     serviceV1.Test,
			ServiceName: "Test",
		},
	}

	// Auth Service
	authServices := []ResponseBaseService{
		{
			MsgType:     constants.MsgTypeAuth,
			MsgAction:   constants.MsgActionAuthLogin,
			Service:     serviceV1.AuthLogin,
			ServiceName: "AuthLogin",
		},
		{
			MsgType:     constants.MsgTypeAuth,
			MsgAction:   constants.MsgActionAuthRegister,
			Service:     serviceV1.AuthRegister,
			ServiceName: "AuthRegister",
		},
	}

	services = append(services, commonServices...)
	services = append(services, authServices...)

	logger.Info("[END] BaseService...")
	return services
}
