package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthcheck responds with 200 if the service is up and running
func Healthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ExecHub API is up and running...")
}
