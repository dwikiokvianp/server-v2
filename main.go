package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"server-v2/config"
	"server-v2/routes"
)

func main() {
	failedLoadEnv := godotenv.Load()
	if failedLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrl := os.Getenv("DB_URL")

	config.InitDatabase(databaseUrl)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my API",
		})
	})
	routes.Routes(r)
	port := os.Getenv("PORT")
	err := r.Run(port)
	if err != nil {
		log.Fatal("Error running server")
	}
}
