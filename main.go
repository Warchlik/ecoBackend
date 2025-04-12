package main

import (
	"eco-backend/app/database"
	"eco-backend/app/routes"
	"eco-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := config.GetEnv("APP_PORT", "8080")
	r.Run(":" + port)
}
