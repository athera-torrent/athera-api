package main

import (
	"athera-api/api/models"
	"athera-api/config"
	"log"
	"os"
	"strings"
	"time"

	"athera-api/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Database
	config.ConnectDB()

	// Auto migrate database models
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.UserClass{},
		&models.Invite{},
		&models.APIAccessKey{},
		&models.TorrentPassKey{},
	); err != nil {
		log.Fatal("Failed to auto migrate database models:", err)
	}

	// Initialize Gin router with Logger and Recovery middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Setup CORS
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	routes.SetupAuthRoutes(router, config.DB)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if not specified
	}

	// Start server
	log.Fatal(router.Run(":" + port))
}
