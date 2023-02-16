package services

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
)

var logLookup = coreLogger.Logger(constants.AppName, constants.StructLookupService)

func LookupService(msgType string, msgAction string) FuncServiceHandler {
	logLookup.Debug("[BEGIN] LookupService with msgType = ", "["+msgType+"]", ", msgAction = ", "["+msgAction+"]")

	services := BaseService()

	var fnServiceHandler FuncServiceHandler

	for _, s := range services {
		if s.MsgType == msgType && s.MsgAction == msgAction {
			fnServiceHandler = s.Service
			logLookup.Debug("END] LookupService Success With Service = ", "["+s.ServiceName+"]")
			return fnServiceHandler
		}
	}

	logLookup.Debug("[END] LookupService Not Found Service With msgType = ", "["+msgType+"]", ", msgAction = ", "["+msgAction+"]")

	return nil
}
