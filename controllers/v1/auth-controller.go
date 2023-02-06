package v1

import (
	constants "fast-food-api-client/constants"
	controllers "fast-food-api-client/controllers"
	coreLogger "fast-food-api-client/core/logger"

	"github.com/gin-gonic/gin"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructAuthController)

func LoginAuth(c *gin.Context) {
	logger.Info("BEGIN LoginAuth Controller")
	controllers.BaseController(c, constants.MsgTypeAuth, constants.MsgActionAuthLogin)
	logger.Info("END LoginAuth Controller")
}
