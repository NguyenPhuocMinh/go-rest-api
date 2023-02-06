package main

import (
	profiles "fast-food-api-client/configs"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	middleware "fast-food-api-client/middleware"
	routerV1 "fast-food-api-client/routers/v1"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructMain)

func main() {
	router := gin.New()
	router.UseRawPath = true
	router.HandleMethodNotAllowed = true
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.NoRoute(middleware.NoRouteMiddleware())
	router.NoMethod(middleware.NoMethodMiddleware())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/api/v1")
	routerV1.InitRouter(v1)

	port := profiles.AppPort
	logger.Info("The server is running on: ", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
