package main

import (
	"github.com/ScotuzziJr/lead-service-exechub/src/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // instantiate gin client

	// routes
	router.GET("/healthcheck", handlers.Healthcheck)
	router.POST("/api/lead", handlers.SaveLead)

	router.Run("localhost:8000")
}
