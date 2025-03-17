package services

import (
	"github.com/yourusername/picto-lingua-backend/api/models"
)

// ThemeService manages themes for the application
type ThemeService struct {
	themes []models.Theme
}

// NewThemeService creates a new theme service with predefined themes
func NewThemeService() *ThemeService {
	return &ThemeService{
		themes: []models.Theme{
			{ID: "cafe", Name: "Café/Coffee Shop", Description: "Vocabulary related to cafés and coffee shops"},
			{ID: "park", Name: "Park/Nature", Description: "Vocabulary related to parks and nature"},
			{ID: "airport", Name: "Airport/Travel", Description: "Vocabulary related to airports and travel"},
			{ID: "kitchen", Name: "Kitchen/Cooking", Description: "Vocabulary related to kitchens and cooking"},
			{ID: "office", Name: "Office/Workplace", Description: "Vocabulary related to offices and workplaces"},
			{ID: "beach", Name: "Beach/Ocean", Description: "Vocabulary related to beaches and oceans"},
			{ID: "city", Name: "City/Urban", Description: "Vocabulary related to cities and urban environments"},
			{ID: "home", Name: "Home/Living Space", Description: "Vocabulary related to homes and living spaces"},
			{ID: "grocery", Name: "Grocery Store/Shopping", Description: "Vocabulary related to grocery stores and shopping"},
			{ID: "restaurant", Name: "Restaurant/Dining", Description: "Vocabulary related to restaurants and dining"},
		},
	}
}

// GetAllThemes returns all available themes
func (s *ThemeService) GetAllThemes() []models.Theme {
	return s.themes
}

// GetThemeByID returns a theme by its ID
func (s *ThemeService) GetThemeByID(id string) *models.Theme {
	for _, theme := range s.themes {
		if theme.ID == id {
			return &theme
		}
	}
	return nil
}

// IsValidTheme checks if a theme ID is valid
func (s *ThemeService) IsValidTheme(id string) bool {
	return s.GetThemeByID(id) != nil
}
