package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/admin"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/user"
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

	adminRoutes.GET("/", func(c *gin.Context) {
		adminController.GetCofig()
	})

	adminRoutes.PATCH("/", func(c *gin.Context) {
		adminController.UpdateConfig()
	})

	adminRoutes.GET("country/:id", func(c *gin.Context) {
		adminController.GetCountryDetails()
	})

	adminRoutes.PUT("country/:id", func(c *gin.Context) {
		adminController.UpdateCountryDetails()
	})

	adminRoutes.POST("country/:id", func(c *gin.Context) {
		adminController.CreateCountryDetails()
	})

	adminRoutes.DELETE("country/:id", func(c *gin.Context) {
		adminController.DeleteCountry()
	})

	app.Run()
}
