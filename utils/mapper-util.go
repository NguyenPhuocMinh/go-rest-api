package utils

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"
)

var logConvert = coreLogger.Logger(constants.AppName, constants.StructMapperUtil)

func MapStatusByMsgCode(msgCode string) int {

	statusCode, ok := resources.MapMsgStatus[msgCode]
	if !ok {
		logConvert.Warn("Cannot found status code with message code = ", msgCode, ", so return default status 500")
		statusCode = constants.StatusInternalServerError
	}

	return statusCode
}

func MapMsgNameByMsgCode(msgCode string) string {

	msgName, ok := resources.MapMsgName[msgCode]
	if !ok {
		logConvert.Warn("Cannot found message name with message code = ", msgCode, ", so return default message name")
		msgName = resources.MapMsgName[resources.MsgCodeNotDefined]
	}

	return msgName
}
