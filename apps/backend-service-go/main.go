package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/user/userController"
)

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[Server] Failed to load environment variables")
	}
}

func main() {
	loadEnvVars()
	app := gin.Default()
	apiRoutes := app.Group("api")
	userRoutes := app.Group("api/user")

	userController := userController.New()

	apiRoutes.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": "OK"})
	})

	userRoutes.GET(":id", func(c *gin.Context) {
		userController.GetUser(c)
	})

	userRoutes.POST("/", func(c *gin.Context) {
		userController.NewUser(c)
	})

	userRoutes.PATCH(":id", func(c *gin.Context) {
		userController.UpdateUser(c)
	})

	userRoutes.DELETE(":id", func(c *gin.Context) {
		userController.DeleteUser(c)
	})

	app.Run()
}
