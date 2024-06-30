package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meetamjadsaeed/task-manager/internal/services"
)

func GenerateReport(c *gin.Context) {
	report, err := services.GenerateReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"report": report})
}
