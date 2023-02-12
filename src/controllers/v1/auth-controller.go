package v1

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	controllers "fast-food-api-client/src/controllers"

	"github.com/gin-gonic/gin"
)

var logAuth = coreLogger.Logger(constants.AppName, constants.StructAuthController)

func LoginAuth(c *gin.Context) {
	logAuth.Info("[BEGIN] LoginAuth Controller...")
	controllers.BaseController(c, constants.MsgTypeAuth, constants.MsgActionAuthLogin)
	logAuth.Info("[END] LoginAuth Controller...")
}

func RegisterAuth(c *gin.Context) {
	logAuth.Info("BEGIN RegisterAuth Controller...")
	controllers.BaseController(c, constants.MsgTypeAuth, constants.MsgActionAuthRegister)
	logAuth.Info("END RegisterAuth Controller...")
}
