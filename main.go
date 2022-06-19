package main

import (
	"main/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(gin.ErrorLogger())
	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	app.Run(":5500")
}
