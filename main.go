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
	routes.Routes(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
