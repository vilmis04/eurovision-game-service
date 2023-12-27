package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[Server] Failed to load environment variables")
	}
}

func main() {
	loadEnvVars()
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": "OK"})
	})

	router.Run()
}
