package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetThemes handles the request to get all available themes
func GetThemes(c *gin.Context) {
	// Get all themes from the service
	themes := themeService.GetAllThemes()

	// Return the themes
	c.JSON(http.StatusOK, gin.H{
		"themes": themes,
	})
}
