package v1

import "github.com/gin-gonic/gin"

func Init(r *gin.RouterGroup) {
	AuthRouter(r)
	CommonRouter(r)
}
