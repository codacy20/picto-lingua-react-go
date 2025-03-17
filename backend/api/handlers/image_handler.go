package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/picto-lingua-backend/api/services"
	"github.com/yourusername/picto-lingua-backend/config"
)

var (
	unsplashService *services.UnsplashService
	themeService    *services.ThemeService
)

// InitImageHandler initializes the image handler with necessary services
func InitImageHandler(cfg *config.Config) {
	unsplashService = services.NewUnsplashService(cfg.UnsplashAccessKey)
	themeService = services.NewThemeService()
}

// GetImages handles the request to get images for a theme
func GetImages(c *gin.Context) {
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

	// Get the count parameter, default to 5
	count := 5

	// Get images from the service
	images, err := unsplashService.SearchImages(theme, count)
	if err != nil {
		log.Printf("Error getting images: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get images"})
		return
	}

	// Return the images
	c.JSON(http.StatusOK, gin.H{
		"theme":  theme,
		"images": images,
	})
}
