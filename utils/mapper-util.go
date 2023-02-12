package utils

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	helpers "fast-food-api-client/helpers"
	resources "fast-food-api-client/resources"
)

var logMapper = coreLogger.Logger(constants.AppName, constants.StructMapperUtil)

func MapStatusByMsgCode(msgCode string) int {
	statusCode, ok := resources.MapMsgStatus[msgCode]
	if !ok {
		logMapper.Warn("Cannot found status code with message code = ", msgCode, ", return default status 500")
		statusCode = constants.StatusInternalServerError
	}

	return statusCode
}

func MapMsgNameByMsgCode(msgCode string) string {
	msgName, ok := resources.MapMsgName[msgCode]
	if !ok {
		logMapper.Warn("Cannot found message name with message code = ", msgCode, ", return default message name")
		msgName = resources.MapMsgName[resources.MsgCodeRequestNotDefined]
	}

	return msgName
}

func MapMsgTranslatorByLangAndMsgCode(lang string, msgCode string) string {
	if helpers.IsNil(lang) {
		lang = constants.LangEN
	}

	msgTranslator, ok := resources.MapMsgTranslator[lang][msgCode]
	if !ok {
		logMapper.Warn("Cannot found message name with message code = ", msgCode, ", return default message translator")
		msgTranslator = resources.MapMsgTranslator[lang][resources.MsgCodeRequestNotDefined]
	}

	return msgTranslator
}
