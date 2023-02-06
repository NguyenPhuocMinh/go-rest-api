package v1

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.RouterGroup) {
	AuthRouter(r)
}
