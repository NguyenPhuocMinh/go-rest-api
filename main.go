package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	http.ListenAndServe(":8080", router)
}
