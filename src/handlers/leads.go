package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/ScotuzziJr/lead-service-exechub/src/config"
	"github.com/ScotuzziJr/lead-service-exechub/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveLead(c *gin.Context) {
	// Parse the request body
	var newLead models.Lead
	if err := c.ShouldBindJSON(&newLead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Generate a new UUID for the lead
	newLead.ID = uuid.New().String()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	// Insert the lead into the database
	query := `INSERT INTO leads (id, name, email, role, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(context.Background(), query, newLead.ID, newLead.Name, newLead.Email, newLead.Role, time.Now())
	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the lead"})
		return
	}

	// Respond with the saved lead
	c.JSON(http.StatusCreated, newLead)
}
