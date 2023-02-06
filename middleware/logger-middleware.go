package middleware

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

var l = coreLogger.Logger(constants.AppName, constants.StructLoggerMiddleware)

func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {

			message := fmt.Sprintf("%s - \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)

			switch param.StatusCode {
			case constants.StatusOK, constants.StatusCreated, constants.StatusAccepted:
				l.Info(message)
			case constants.StatusBadRequest, constants.StatusUnauthorized, constants.StatusForbidden, constants.StatusNotFound, constants.StatusConflict:
				l.Warn(message)
			case constants.StatusMethodNotAllowed:
				l.Debug(message)
			case constants.StatusInternalServerError:
				l.Error(message)
			}

			return ""
		},
	})
}
