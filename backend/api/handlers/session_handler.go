package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/picto-lingua-backend/api/models"
	"github.com/yourusername/picto-lingua-backend/api/services"
)

var (
	sessionService *services.SessionService
)

// InitSessionHandler initializes the session handler with necessary services
func InitSessionHandler() {
	sessionService = services.NewSessionService()
}

// SaveSession handles the request to save a user's session data
func SaveSession(c *gin.Context) {
	// Parse the session data from the request body
	var sessionRequest struct {
		ThemeID   string                         `json:"theme_id" binding:"required"`
		ImageID   string                         `json:"image_id" binding:"required"`
		Progress  map[string]models.ProgressItem `json:"progress"`
		SessionID string                         `json:"session_id"`
	}

	if err := c.ShouldBindJSON(&sessionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the theme
	if !themeService.IsValidTheme(sessionRequest.ThemeID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid theme"})
		return
	}

	var sessionID string
	var err error

	// If no session ID is provided, create a new session
	if sessionRequest.SessionID == "" {
		sessionID, err = sessionService.CreateSession(sessionRequest.ThemeID, sessionRequest.ImageID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
			return
		}
	} else {
		// Otherwise update the existing session
		sessionID = sessionRequest.SessionID
		err = sessionService.UpdateSession(sessionID, sessionRequest.Progress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update session"})
			return
		}
	}

	// Return the session ID
	c.JSON(http.StatusOK, gin.H{
		"session_id": sessionID,
		"status":     "success",
	})
}

// GetSession handles the request to get a user's session data
func GetSession(c *gin.Context) {
	// Get the session ID from the query parameters
	sessionID := c.Query("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	// Get the session from the service
	session, err := sessionService.GetSession(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	// Return the session data
	c.JSON(http.StatusOK, session)
}
