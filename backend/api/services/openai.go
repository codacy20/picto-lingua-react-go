package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/yourusername/picto-lingua-backend/api/models"
)

// OpenAIService handles communication with the OpenAI API
type OpenAIService struct {
	client *openai.Client
}

// NewOpenAIService creates a new OpenAI service
func NewOpenAIService(apiKey string) *OpenAIService {
	client := openai.NewClient(apiKey)
	return &OpenAIService{
		client: client,
	}
}

// GenerateVocabulary generates vocabulary words for a given theme
func (s *OpenAIService) GenerateVocabulary(theme string, count int) ([]models.VocabularyItem, error) {
	prompt := fmt.Sprintf(`Generate %d vocabulary words related to the theme "%s". 
Each word should have a definition and a simple example sentence.
Format your response as a JSON array of objects, where each object contains:
- "word": the vocabulary word
- "definition": a brief definition of the word
- "example": a simple example sentence using the word

Only provide the JSON output, no additional text.`, count, theme)

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
		return nil, fmt.Errorf("error generating vocabulary: %w", err)
	}

	// Parse the JSON response
	var vocabulary []models.VocabularyItem
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &vocabulary); err != nil {
		return nil, fmt.Errorf("error parsing vocabulary response: %w", err)
	}

	return vocabulary, nil
}

// Cache to store previously generated vocabulary
var vocabularyCache = make(map[string][]models.VocabularyItem)

// GetVocabularyWithCache gets vocabulary for a theme using caching
func (s *OpenAIService) GetVocabularyWithCache(theme string, count int) ([]models.VocabularyItem, error) {
	cacheKey := fmt.Sprintf("%s_%d", theme, count)

	// Check if we have cached results
	if cachedVocab, ok := vocabularyCache[cacheKey]; ok {
		return cachedVocab, nil
	}

	// Generate new vocabulary
	vocabulary, err := s.GenerateVocabulary(theme, count)
	if err != nil {
		return nil, err
	}

	// Cache the results
	vocabularyCache[cacheKey] = vocabulary

	return vocabulary, nil
}
