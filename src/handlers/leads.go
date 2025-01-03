package handlers

import (
	"net/http"
	"time"

	"github.com/ScotuzziJr/lead-service-exechub/src/config"
	"github.com/ScotuzziJr/lead-service-exechub/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SaveLead handles the saving of a lead to the database.
func SaveLead(c *gin.Context) {
	// Parse the request body
	var newLead models.Lead
	if err := c.ShouldBindJSON(&newLead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Generate a new UUID for the lead
	newLead.ID = uuid.New().String()
	newLead.CreatedAt = time.Now()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	// Save the lead to the database using GORM
	if err := db.Create(&newLead).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the lead"})
		return
	}

	// Respond with the saved lead
	c.JSON(http.StatusCreated, newLead)
}
