package services

import (
	"errors"
	"sync"
	"time"

	"github.com/yourusername/picto-lingua-backend/api/models"
)

// SessionService manages user sessions
type SessionService struct {
	sessions map[string]models.SessionData
	mu       sync.RWMutex
}

// NewSessionService creates a new session service
func NewSessionService() *SessionService {
	return &SessionService{
		sessions: make(map[string]models.SessionData),
	}
}

// CreateSession creates a new session for a user
func (s *SessionService) CreateSession(themeID, imageID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate a simple session ID (in production, use a proper UUID generator)
	sessionID := generateSessionID()

	now := time.Now().Format(time.RFC3339)

	// Create a new session
	session := models.SessionData{
		ThemeID:     themeID,
		ImageID:     imageID,
		Progress:    make(map[string]models.ProgressItem),
		SessionID:   sessionID,
		StartedAt:   now,
		LastUpdated: now,
	}

	// Store the session
	s.sessions[sessionID] = session

	return sessionID, nil
}

// GetSession gets a session by its ID
func (s *SessionService) GetSession(sessionID string) (*models.SessionData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.sessions[sessionID]
	if !ok {
		return nil, errors.New("session not found")
	}

	return &session, nil
}

// UpdateSession updates a session with new progress
func (s *SessionService) UpdateSession(sessionID string, progress map[string]models.ProgressItem) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, ok := s.sessions[sessionID]
	if !ok {
		return errors.New("session not found")
	}

	// Update the progress
	for word, item := range progress {
		session.Progress[word] = item
	}

	// Update the last updated timestamp
	session.LastUpdated = time.Now().Format(time.RFC3339)

	// Store the updated session
	s.sessions[sessionID] = session

	return nil
}

// generateSessionID generates a simple session ID
// In production, use a proper UUID generator
func generateSessionID() string {
	return time.Now().Format("20060102150405.000000000")
}
