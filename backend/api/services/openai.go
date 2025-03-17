package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/yourusername/picto-lingua-backend/api/models"
)

var (
	// Create a logger that writes to a file
	debugLogger = initDebugLogger()
)

// initDebugLogger initializes a logger that writes to a file
func initDebugLogger() *log.Logger {
	// Create or open the log file
	logFile, err := os.OpenFile("openai_debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("ERROR: Failed to open log file: %v", err)
		return log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime)
	}
	return log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime)
}

// OpenAIService handles communication with the OpenAI API
type OpenAIService struct {
	client     *openai.Client
	useMock    bool
	mockThemes map[string][]models.VocabularyItem
}

// NewOpenAIService creates a new OpenAI service
func NewOpenAIService(apiKey string) *OpenAIService {
	service := &OpenAIService{
		mockThemes: make(map[string][]models.VocabularyItem),
	}

	// Check if API key is provided
	if apiKey == "" {
		debugLogger.Printf("WARNING: No OpenAI API key provided, using mock implementation")
		service.useMock = true
		// Initialize mock data
		service.initMockData()
		return service
	}

	debugLogger.Printf("Initializing OpenAI service with API key: %s...", apiKey[:5]+"...")
	client := openai.NewClient(apiKey)
	service.client = client
	return service
}

// initMockData initializes mock vocabulary data for testing
func (s *OpenAIService) initMockData() {
	// Mock data for park theme
	s.mockThemes["park"] = []models.VocabularyItem{
		{Word: "bench", Definition: "A long seat for two or more people", Example: "We sat on the bench in the park."},
		{Word: "playground", Definition: "An area for children with swings, slides, etc.", Example: "The children had fun at the playground."},
		{Word: "fountain", Definition: "An ornamental structure that sends water into the air", Example: "The fountain in the park was beautiful."},
		{Word: "path", Definition: "A way or track for walking or cycling", Example: "We walked along the path through the park."},
		{Word: "tree", Definition: "A tall plant with a wooden trunk and branches", Example: "The trees in the park provide shade in summer."},
		{Word: "grass", Definition: "Plants with narrow green leaves that cover the ground", Example: "The grass in the park was freshly cut."},
		{Word: "picnic", Definition: "An outdoor meal", Example: "We had a picnic in the park on Sunday."},
		{Word: "jogger", Definition: "A person who runs at a steady speed for exercise", Example: "Joggers often use the park in the morning."},
		{Word: "lake", Definition: "A large area of water surrounded by land", Example: "There is a small lake in the center of the park."},
		{Word: "garden", Definition: "An area where flowers and plants are grown", Example: "The botanical garden in the park has rare flowers."},
	}

	// Mock data for cafe theme
	s.mockThemes["cafe"] = []models.VocabularyItem{
		{Word: "coffee", Definition: "A hot drink made from roasted coffee beans", Example: "I ordered a coffee at the cafe."},
		{Word: "barista", Definition: "A person who makes and serves coffee", Example: "The barista made a beautiful design in my latte."},
		{Word: "menu", Definition: "A list of food and drinks available", Example: "The cafe has a varied menu with many options."},
		{Word: "pastry", Definition: "A sweet baked food made with dough", Example: "The cafe sells delicious pastries."},
		{Word: "table", Definition: "A piece of furniture with a flat top", Example: "We found a table by the window in the cafe."},
		{Word: "espresso", Definition: "A strong coffee made by forcing steam through ground coffee beans", Example: "An espresso is perfect for a quick caffeine boost."},
		{Word: "latte", Definition: "Coffee made with hot milk", Example: "She ordered a vanilla latte at the cafe."},
		{Word: "wifi", Definition: "Wireless internet connection", Example: "The cafe offers free wifi to customers."},
		{Word: "ambiance", Definition: "The character and atmosphere of a place", Example: "The cafe has a cozy ambiance with soft lighting."},
		{Word: "tip", Definition: "Money given to a server as a reward for good service", Example: "I left a generous tip at the cafe."},
	}

	// Add more mock themes as needed
}

// GenerateVocabulary generates vocabulary words for a given theme
func (s *OpenAIService) GenerateVocabulary(theme string, count int) ([]models.VocabularyItem, error) {
	debugLogger.Printf("Generating vocabulary for theme: %s, count: %d", theme, count)

	// If using mock implementation, return mock data
	if s.useMock {
		debugLogger.Printf("Using mock implementation for theme: %s", theme)
		mockData, ok := s.mockThemes[theme]
		if !ok {
			return nil, fmt.Errorf("mock data not available for theme: %s", theme)
		}

		// Return the requested number of items, or all items if count > available items
		resultCount := count
		if resultCount > len(mockData) {
			resultCount = len(mockData)
		}

		return mockData[:resultCount], nil
	}

	// Check if client is initialized
	if s.client == nil {
		return nil, fmt.Errorf("OpenAI client not initialized")
	}

	prompt := fmt.Sprintf(`Generate %d vocabulary words related to the theme "%s". 
Each word should have a definition and a simple example sentence.
Format your response as a JSON array of objects, where each object contains:
- "word": the vocabulary word
- "definition": a brief definition of the word
- "example": a simple example sentence using the word

Only provide the JSON output, no additional text.`, count, theme)

	debugLogger.Printf("Using prompt: %s", prompt)

	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a language learning tool that generates vocabulary words with definitions and examples.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.7,
		},
	)

	if err != nil {
		debugLogger.Printf("Error generating vocabulary: %v", err)
		return nil, fmt.Errorf("error generating vocabulary: %w", err)
	}

	// Log the raw response for debugging
	rawResponse := resp.Choices[0].Message.Content
	debugLogger.Printf("Raw OpenAI response: %s", rawResponse)

	// Clean the response if it contains markdown code blocks
	cleanResponse := rawResponse
	if strings.Contains(rawResponse, "```json") {
		cleanResponse = strings.ReplaceAll(rawResponse, "```json", "")
		cleanResponse = strings.ReplaceAll(cleanResponse, "```", "")
		cleanResponse = strings.TrimSpace(cleanResponse)
		debugLogger.Printf("Cleaned JSON response: %s", cleanResponse)
	}

	// Parse the JSON response
	var vocabulary []models.VocabularyItem
	if err := json.Unmarshal([]byte(cleanResponse), &vocabulary); err != nil {
		// Log more details about parsing error
		debugLogger.Printf("JSON parsing error: %v for content: %s", err, cleanResponse)
		return nil, fmt.Errorf("error parsing vocabulary response: %w", err)
	}

	debugLogger.Printf("Successfully parsed %d vocabulary items", len(vocabulary))
	return vocabulary, nil
}

// Cache to store previously generated vocabulary
var vocabularyCache = make(map[string][]models.VocabularyItem)

// GetVocabularyWithCache gets vocabulary for a theme using caching
func (s *OpenAIService) GetVocabularyWithCache(theme string, count int) ([]models.VocabularyItem, error) {
	cacheKey := fmt.Sprintf("%s_%d", theme, count)
	debugLogger.Printf("Getting vocabulary for cache key: %s", cacheKey)

	// Check if we have cached results
	if cachedVocab, ok := vocabularyCache[cacheKey]; ok {
		debugLogger.Printf("Cache hit for key: %s, returning %d vocabulary items", cacheKey, len(cachedVocab))
		return cachedVocab, nil
	}

	debugLogger.Printf("Cache miss for key: %s, generating new vocabulary", cacheKey)
	// Generate new vocabulary
	vocabulary, err := s.GenerateVocabulary(theme, count)
	if err != nil {
		debugLogger.Printf("Error generating vocabulary: %v", err)
		return nil, err
	}

	// Cache the results
	vocabularyCache[cacheKey] = vocabulary
	debugLogger.Printf("Cached %d vocabulary items for key: %s", len(vocabulary), cacheKey)

	return vocabulary, nil
}
