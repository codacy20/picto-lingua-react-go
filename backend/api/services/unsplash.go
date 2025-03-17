package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/yourusername/picto-lingua-backend/api/models"
)

const (
	unsplashBaseURL = "https://api.unsplash.com"
)

// UnsplashService handles communication with the Unsplash API
type UnsplashService struct {
	apiKey string
	client *http.Client
}

// NewUnsplashService creates a new Unsplash service
func NewUnsplashService(apiKey string) *UnsplashService {
	return &UnsplashService{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

// SearchImages searches for images based on a query
func (s *UnsplashService) SearchImages(query string, count int) ([]models.Image, error) {
	endpoint := fmt.Sprintf("%s/search/photos", unsplashBaseURL)

	// Build the URL with query parameters
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	q := u.Query()
	q.Set("query", query)
	q.Set("per_page", fmt.Sprintf("%d", count))
	q.Set("orientation", "landscape") // Prefer landscape for better display
	u.RawQuery = q.Encode()

	// Create the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", s.apiKey))

	// Perform the request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response
	var searchResponse struct {
		Results []struct {
			ID          string `json:"id"`
			Description string `json:"description"`
			Width       int    `json:"width"`
			Height      int    `json:"height"`
			CreatedAt   string `json:"created_at"`
			URLs        struct {
				Raw     string `json:"raw"`
				Regular string `json:"regular"`
				Small   string `json:"small"`
			} `json:"urls"`
			Links struct {
				Download string `json:"download"`
				HTML     string `json:"html"`
			} `json:"links"`
			User struct {
				Name         string `json:"name"`
				PortfolioURL string `json:"portfolio_url"`
				Links        struct {
					HTML string `json:"html"`
				} `json:"links"`
			} `json:"user"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	// Map the response to our model
	images := make([]models.Image, 0, len(searchResponse.Results))
	for _, result := range searchResponse.Results {
		// Apply dynamic resizing parameters
		resizedURL := result.URLs.Regular + "&w=800&h=600&fit=crop&crop=entropy"

		// Create attribution string
		attribution := fmt.Sprintf("Photo by %s on Unsplash", result.User.Name)

		images = append(images, models.Image{
			ID:                result.ID,
			URL:               resizedURL,
			DownloadURL:       result.Links.Download,
			Description:       result.Description,
			Width:             result.Width,
			Height:            result.Height,
			CreatedAt:         result.CreatedAt,
			Photographer:      result.User.Name,
			PhotographerURL:   result.User.Links.HTML,
			UnsplashURL:       result.Links.HTML,
			AttributionString: attribution,
		})
	}

	return images, nil
}

// GetRandomImage gets a random image based on a theme
func (s *UnsplashService) GetRandomImage(theme string) (*models.Image, error) {
	endpoint := fmt.Sprintf("%s/photos/random", unsplashBaseURL)

	// Build the URL with query parameters
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	q := u.Query()
	q.Set("query", theme)
	q.Set("orientation", "landscape")
	u.RawQuery = q.Encode()

	// Create the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", s.apiKey))

	// Perform the request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response
	var result struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
		CreatedAt   string `json:"created_at"`
		URLs        struct {
			Raw     string `json:"raw"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
		} `json:"urls"`
		Links struct {
			Download string `json:"download"`
			HTML     string `json:"html"`
		} `json:"links"`
		User struct {
			Name         string `json:"name"`
			PortfolioURL string `json:"portfolio_url"`
			Links        struct {
				HTML string `json:"html"`
			} `json:"links"`
		} `json:"user"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	// Apply dynamic resizing parameters
	resizedURL := result.URLs.Regular + "&w=800&h=600&fit=crop&crop=entropy"

	// Create attribution string
	attribution := fmt.Sprintf("Photo by %s on Unsplash", result.User.Name)

	image := &models.Image{
		ID:                result.ID,
		URL:               resizedURL,
		DownloadURL:       result.Links.Download,
		Description:       result.Description,
		Width:             result.Width,
		Height:            result.Height,
		CreatedAt:         result.CreatedAt,
		Photographer:      result.User.Name,
		PhotographerURL:   result.User.Links.HTML,
		UnsplashURL:       result.Links.HTML,
		AttributionString: attribution,
	}

	return image, nil
}
