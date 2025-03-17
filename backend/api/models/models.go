package models

// Image represents an image from Unsplash
type Image struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	DownloadURL string `json:"download_url"`
	Description string `json:"description"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	CreatedAt   string `json:"created_at"`
	// Attribution information
	Photographer      string `json:"photographer"`
	PhotographerURL   string `json:"photographer_url"`
	UnsplashURL       string `json:"unsplash_url"`
	AttributionString string `json:"attribution_string"`
}

// VocabularyItem represents a vocabulary word and its definition
type VocabularyItem struct {
	Word            string `json:"word"`
	Definition      string `json:"definition"`
	Example         string `json:"example,omitempty"`
	DutchWord       string `json:"dutch_word,omitempty"`
	DutchDefinition string `json:"dutch_definition,omitempty"`
	DutchExample    string `json:"dutch_example,omitempty"`
}

// Theme represents a learning theme
type Theme struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// SessionData represents a user's learning session data
type SessionData struct {
	ThemeID     string                  `json:"theme_id"`
	ImageID     string                  `json:"image_id"`
	Progress    map[string]ProgressItem `json:"progress"`
	SessionID   string                  `json:"session_id"`
	StartedAt   string                  `json:"started_at"`
	LastUpdated string                  `json:"last_updated"`
}

// ProgressItem represents a user's progress on a specific vocabulary item
type ProgressItem struct {
	Word       string `json:"word"`
	Status     string `json:"status"` // "known", "learning", "difficult"
	TimeTaken  int    `json:"time_taken_ms,omitempty"`
	SeenCount  int    `json:"seen_count"`
	KnownCount int    `json:"known_count"`
}
