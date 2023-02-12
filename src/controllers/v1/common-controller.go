package v1

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	controllers "fast-food-api-client/src/controllers"

	"github.com/gin-gonic/gin"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructCommonController)

func Test(c *gin.Context) {
	logger.Info("[BEGIN] Test Controller")
	controllers.BaseController(c, constants.MsgTypeCommon, constants.MsgActionTest)
	logger.Info("[END] Test Controller")
}
