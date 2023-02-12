package main

import (
	"fmt"
	"net/http"

	configs "fast-food-api-client/configs"
	constants "fast-food-api-client/constants"
	database "fast-food-api-client/core/database"
	coreLogger "fast-food-api-client/core/logger"
	middleware "fast-food-api-client/middleware"
	routerV1 "fast-food-api-client/src/routers/v1"

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

	// init router
	v1 := router.Group("/api/v1")
	routerV1.Init(v1)

	// Init db
	mongoURI := configs.AppMongoURI
	database.Init(mongoURI)

	port := configs.AppPort
	logger.Info("The server is running on: ", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
