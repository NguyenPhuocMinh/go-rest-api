package middleware

import (
	"errors"
	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"

	"github.com/gin-gonic/gin"
)

var logNoRoute = coreLogger.Logger(constants.AppName, constants.StructNoRouteMiddleware)

func NoRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logNoRoute.Debug("BEGIN NoRouteMiddleware...")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(constants.StatusNoContent)
			return
		}

		err := errors.New("Route not found for request path: " + c.Request.URL.String())

		logNoRoute.Debug("END NoRouteMiddleware with err= ", err.Error())
		res := commons.TemplateErrorCommon(c, err, resources.MsgCodeRouteNotFound)
		res.ResponseCommon(c)
	}
}
