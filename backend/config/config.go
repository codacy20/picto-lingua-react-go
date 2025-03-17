package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	UnsplashAccessKey string
	OpenAIAPIKey      string
	Port              string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{
		UnsplashAccessKey: getEnv("UNSPLASH_ACCESS_KEY", ""),
		OpenAIAPIKey:      getEnv("OPENAI_API_KEY", ""),
		Port:              getEnv("PORT", "8080"),
	}

	return config, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
