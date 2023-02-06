package v1

import (
	controllerV1 "fast-food-api-client/controllers/v1"
	middleware "fast-food-api-client/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	rAuth := r.Group("/auth")
	{
		rAuth.POST("/login", middleware.AuthMiddleware(), controllerV1.LoginAuth)
	}
}
