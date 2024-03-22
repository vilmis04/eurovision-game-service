package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/country"
	"github.com/vilmis04/eurovision-game-service/internal/group"
	"github.com/vilmis04/eurovision-game-service/internal/user"
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
	app := gin.Default()
	apiRoutes := app.Group("api")

	apiRoutes.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": "OK"})
	})

	user.NewController(app).Use()
	admin.NewController(app).Use()
	country.NewController(app).Use()
	group.NewController(app).Use()

	app.Run()
}
