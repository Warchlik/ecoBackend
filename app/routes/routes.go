package routes

import (
	// "eco-backend/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		// api.POST("/register", controllers.Register)
		// api.POST("/login", controllers.Login)
	}
}
