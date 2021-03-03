package main

import (
	"base64/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	controllers.RegisterRoutes(router)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router.LoadHTMLGlob("templates/*")
	router.Static("/css", "static/css")
	router.Run()
}
