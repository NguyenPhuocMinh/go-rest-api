package middleware

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"

	"github.com/gin-gonic/gin"
)

var logAuth = coreLogger.Logger(constants.AppName, constants.StructAuthMiddleware)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logAuth.Debug("BEGIN authMiddleware...")
		c.Next()
	}
}
