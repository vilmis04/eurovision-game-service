package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/admin"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/user"
)

func loadEnvVars() {
	PORT := os.Getenv("PORT")
	fmt.Printf("the port: %v!", PORT)
	if PORT == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("[Server] Failed to load environment variables")
		}
	}
}

func init() {
	loadEnvVars()
}

func main() {
	// loadEnvVars()
	app := gin.Default()
	apiRoutes := app.Group("api")
	userRoutes := app.Group("api/user")
	adminRoutes := app.Group("api/admin")

	userController := user.NewController()
	adminController := admin.NewController()

	apiRoutes.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": "OK"})
	})

	// user routes

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

	// admin routes
	// TODO: add auth middleware for admin endpoints
	adminRoutes.GET("/", func(c *gin.Context) {
		adminController.GetConfig(c)
	})

	adminRoutes.PATCH("/", func(c *gin.Context) {
		adminController.UpdateConfig(c)
	})

	adminRoutes.GET("country/:id", func(c *gin.Context) {
		adminController.GetCountryDetails(c)
	})

	adminRoutes.PUT("country/:id", func(c *gin.Context) {
		adminController.UpdateCountryDetails(c)
	})

	adminRoutes.POST("country/:id", func(c *gin.Context) {
		adminController.CreateCountryDetails(c)
	})

	adminRoutes.DELETE("country/:id", func(c *gin.Context) {
		adminController.DeleteCountry(c)
	})

	app.Run()
}
