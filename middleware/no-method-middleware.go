package middleware

import (
	"errors"

	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"

	"github.com/gin-gonic/gin"
)

var logNoMethod = coreLogger.Logger(constants.AppName, constants.StructNoMethodMiddleware)

func NoMethodMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logNoMethod.Debug("BEGIN NoMethodMiddleware...")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(constants.StatusNoContent)
			return
		}

		err := errors.New("Method not found for request path: " + c.Request.URL.String())

		logNoMethod.Info("END NoMethodMiddleware with err= ", err.Error())
		res := commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestMethodNotFound)
		res.ResponseCommon(c)
	}
}
