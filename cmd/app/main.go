package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/auth"
	"github.com/vilmis04/eurovision-game-service/internal/country"
	"github.com/vilmis04/eurovision-game-service/internal/group"
	"github.com/vilmis04/eurovision-game-service/internal/score"
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

// requireEnv stops the service at startup instead of running without a secret.
func requireEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("[Server] %v must be set", name)
	}

	return value
}

func init() {
	loadEnvVars()
}

func main() {
	internalToken := requireEnv("INTERNAL_TOKEN")
	requireEnv("INVITE_SECRET")

	app := gin.Default()

	// health is registered before the proxy middleware so it stays unauthenticated
	app.GET("api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": "OK"})
	})

	// every route registered below trusts the `user` header only when the
	// request carries the shared token that the auth proxy adds
	app.Use(auth.Proxy(internalToken))

	admin.NewController(app).Use()
	country.NewController(app).Use()
	group.NewController(app).Use()
	score.NewController(app).Use()

	app.Run()
}
