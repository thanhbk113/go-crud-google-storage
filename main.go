package main

import (
	"my-app/config"
	"my-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	port := config.EnvGet("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	routes.CloudRoutes(router)

	router.Run(":" + port)
}
