package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/picto-lingua-backend/api/handlers"
	"github.com/yourusername/picto-lingua-backend/config"
)

func main() {
	// Set up configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize handlers with services
	handlers.InitImageHandler(cfg)
	handlers.InitVocabularyHandler(cfg)
	handlers.InitSessionHandler()

	// Set up the router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Set up the API routes
	api := router.Group("/api")
	{
		// Image routes
		api.GET("/images", handlers.GetImages)

		// Vocabulary routes
		api.GET("/vocabulary", handlers.GetVocabulary)

		// Session routes
		api.POST("/session", handlers.SaveSession)
		api.GET("/session", handlers.GetSession)

		// Theme routes
		api.GET("/themes", handlers.GetThemes)
	}

	// Start the server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
