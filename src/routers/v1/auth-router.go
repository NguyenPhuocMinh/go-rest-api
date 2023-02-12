package v1

import (
	middleware "fast-food-api-client/middleware"
	controllerV1 "fast-food-api-client/src/controllers/v1"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	rAuth := r.Group("/auth", middleware.AuthMiddleware())
	{
		rAuth.POST("/login", controllerV1.LoginAuth)
		rAuth.POST("/register", controllerV1.RegisterAuth)
	}
}
