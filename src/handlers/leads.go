package handlers

import (
	"log"
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
    
    // pegando o body da requisição
    if err := c.ShouldBindJSON(&newLead); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body","msg": err})
        return
    }

    // user := models.Lead{ID: uuid.New().String(), Name: "Wellington", Email: "wellingtons.bezerra@hotmail.com", Role: 0, CreatedAt: time.Now().UTC()}


    // Generate a new UUID for the lead
    newLead.ID = uuid.New().String()
    newLead.CreatedAt = time.Now().Local().UTC()

    // Connect to the database
    db, err := config.ConnectDB()

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
        return
    }

    // Save the lead to the database using GORM
	
    result := db.Create(&newLead)
    if result.Error != nil {
        log.Printf("Error inserting lead: %v", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return // Important: return here to stop further execution
    }

    response := struct {
        RowsAffected int64       `json:"rows_affected"`
        Data         models.Lead `json:"data"`
    }{
        RowsAffected: result.RowsAffected,
        Data:         newLead,
    }

    // Respond with the saved lead
    c.JSON(http.StatusCreated, response)
}
