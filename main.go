package main

import (
	"log"

	"github.com/dheepssupreme/project.git/config"
	"github.com/dheepssupreme/project.git/database"
	"github.com/dheepssupreme/project.git/handlers"
	"github.com/dheepssupreme/project.git/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	database.ConnectDatabase(cfg.DatabaseURL)

	// Setup Gin router
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://yourdomain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Server is running"})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		// Public routes
		public := api.Group("/")
		public.Use(middleware.OptionalAuthMiddleware(cfg.JWTSecret))
		{
			// Articles
			public.GET("/articles", handlers.GetArticles)
			public.GET("/articles/:slug", handlers.GetArticle)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			// Articles (CRUD)
			protected.POST("/articles", handlers.CreateArticle)
			protected.PUT("/articles/:slug", handlers.UpdateArticle)
			protected.DELETE("/articles/:slug", handlers.DeleteArticle)
		}
	}

	log.Printf("Server starting on port %s", cfg.Port)
	log.Printf("Environment: %s", cfg.Environment)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 