package v1

import (
	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	"fast-food-api-client/resources"

	"github.com/gin-gonic/gin"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructAuthService)

func AuthLogin(c *gin.Context) *commons.ResponseModel {
	logger.Info("BEGIN AuthLogin Service")
	res := commons.TemplateSuccessCommon(c, "hehe", resources.MsgCodeRequestDataSuccess)
	logger.Info("END AuthLogin Service")
	return res
}
