package controllers

import (
	"errors"

	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"
	services "fast-food-api-client/src/services"

	"github.com/gin-gonic/gin"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructController)

func BaseController(c *gin.Context, msgType string, msgAction string) {
	logger.Info("[BEGIN] BaseController with msgType = ", "["+msgType+"]", ", msgAction = ", "["+msgAction+"]")
	// lookup service with type and action
	fnServiceHandler := services.LookupService(msgType, msgAction)

	if fnServiceHandler == nil {
		logger.Error("Not Found Service Handler")
		err := errors.New("Service handler not found with msgType = " + "[" + msgType + "]" + ", " + "msgAction = " + "[" + msgAction + "]")
		res := commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestServiceNotFound)
		res.ResponseCommon(c)
	}
	// return response from service handler
	res := fnServiceHandler(c)
	// render json response common
	logger.Info("[END] BaseController with msgType = ", "["+msgType+"]", ", msgAction = ", "["+msgAction+"]")
	res.ResponseCommon(c)
}
