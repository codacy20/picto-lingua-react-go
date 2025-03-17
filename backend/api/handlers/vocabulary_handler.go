package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/picto-lingua-backend/api/services"
	"github.com/yourusername/picto-lingua-backend/config"
)

var (
	openAIService *services.OpenAIService
)

// InitVocabularyHandler initializes the vocabulary handler with necessary services
func InitVocabularyHandler(cfg *config.Config) {
	openAIService = services.NewOpenAIService(cfg.OpenAIAPIKey)
}

// GetVocabulary handles the request to get vocabulary for a theme
func GetVocabulary(c *gin.Context) {
	// Get the theme from the query parameters
	theme := c.Query("theme")
	if theme == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "theme is required"})
		return
	}

	// Validate the theme
	if !themeService.IsValidTheme(theme) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid theme"})
		return
	}

	// Get the count parameter, default to 10
	countStr := c.DefaultQuery("count", "10")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	// Limit count to reasonable bounds
	if count < 1 {
		count = 1
	} else if count > 20 {
		count = 20
	}

	// Get vocabulary from the service (with caching)
	vocabulary, err := openAIService.GetVocabularyWithCache(theme, count)
	if err != nil {
		log.Printf("Error getting vocabulary: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get vocabulary"})
		return
	}

	// Return the vocabulary
	c.JSON(http.StatusOK, gin.H{
		"theme":      theme,
		"count":      len(vocabulary),
		"vocabulary": vocabulary,
	})
}
