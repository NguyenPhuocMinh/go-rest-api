package services

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
)

var logLookup = coreLogger.Logger(constants.AppName, constants.StructLookupService)

func LookupService(msgType string, msgAction string) FuncServiceHandler {
	logLookup.Info("BEGIN LookupService with msgType=", "["+msgType+"]", ", msgAction=", "["+msgAction+"]")

	services := BaseService()

	var fnServiceHandler FuncServiceHandler

	for _, s := range services {
		if s.MsgType == msgType && s.MsgAction == msgAction {
			fnServiceHandler = s.Service
			logLookup.Info("Lookup service success with service=", s.ServiceName)
			return fnServiceHandler
		}
	}

	logLookup.Info("END LookupService with msgType=", "["+msgType+"]", ", msgAction=", "["+msgAction+"]")

	return fnServiceHandler
}
