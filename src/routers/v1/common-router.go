package v1

import (
	middleware "fast-food-api-client/middleware"
	controllerV1 "fast-food-api-client/src/controllers/v1"

	"github.com/gin-gonic/gin"
)

func CommonRouter(r *gin.RouterGroup) {
	rCommon := r.Group("/commons", middleware.AuthMiddleware())
	{
		rCommon.POST("/tests", controllerV1.Test)
	}
}
