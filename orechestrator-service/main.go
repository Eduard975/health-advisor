package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"health-harbor-backend/database"
	"health-harbor-backend/handlers"
	"health-harbor-backend/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	router := gin.Default()

	router.Use(middleware.CORS())

	// the only public routes
	router.POST("/api/auth/login", handlers.Login)
	router.POST("/api/auth/register", handlers.Register)

	// the protected routse
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
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}