package main

import (
	"fmt"
	"time"

	"github.com/ScotuzziJr/lead-service-exechub/src/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env") // load .env

	if err != nil {
		fmt.Errorf("error loading .env file: %w", err)
	}

	router := gin.Default() // instantiate gin client

	router.Use(
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:5173"},
			AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
			AllowHeaders: []string{"Origin", "Content-Type"},
			ExposeHeaders: []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool{
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		}))
	// routes
	router.GET("/healthcheck", handlers.Healthcheck)
	router.POST("/api/lead", handlers.SaveLead)

	router.Run("localhost:8000")
}
