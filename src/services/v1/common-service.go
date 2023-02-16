package v1

import (
	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"

	"github.com/gin-gonic/gin"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructCommonervice)

func Test(c *gin.Context) *commons.ResponseModel {
	logger.Info("[BEGIN] Test Service")

	var result *commons.ResponseModel

	result = commons.TemplateSuccessCommon(c, "test", resources.MsgCodeRequestDataSuccess)
	logger.Info("[END] Test Service")
	return result
}
