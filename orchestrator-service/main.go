package main

import (
	"log"
	"os"

	"orchestrator-service/database"
	"orchestrator-service/handlers"
	"orchestrator-service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.InitFirebase()
	defer database.CloseFirebase()

	router := gin.Default()

	router.Use(middleware.CORS())

	// Auth routes
	router.POST("/api/auth/login", handlers.Login)
	router.POST("/api/auth/register", handlers.Register)
	router.POST("/api/auth/google", handlers.GoogleAuth)

	// Protected routes
	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", handlers.GetProfile)
		auth.PUT("/profile", handlers.UpdateProfile)

		auth.GET("/activity", handlers.GetActivity)
		auth.POST("/activity", handlers.CreateActivity)
		auth.GET("/activity/summary", handlers.GetActivitySummary)

		auth.GET("/health-records", handlers.GetHealthRecords)
		auth.POST("/health-records", handlers.CreateHealthRecord)
		auth.DELETE("/health-records/:id", handlers.DeleteHealthRecord)

		auth.POST("/chat", handlers.SendMessage)
		auth.GET("/chat/history", handlers.GetChatHistory)

		auth.PUT("/settings", handlers.UpdateSettings)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
