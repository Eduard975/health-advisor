package handlers

import (
	"context"
	"net/http"
	"time"

	"health-harbor-backend/database"
	"health-harbor-backend/models"
	"health-harbor-backend/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	iter := database.Client.Collection("users").Where("email", "==", req.Email).Limit(1).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		if err == iterator.Done {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// For Google users, don't allow password login
	if user.Provider == "google" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please use Google Sign-In"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"fullName": user.FullName,
		},
	})
}

func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// Check if user already exists
	iter := database.Client.Collection("users").Where("email", "==", req.Email).Limit(1).Documents(ctx)
	defer iter.Stop()

	if _, err := iter.Next(); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create user
	user := models.User{
		ID:        utils.GenerateID(),
		Email:     req.Email,
		Password:  hashedPassword,
		FullName:  req.FullName,
		Provider:  "email",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Settings: models.UserSettings{
			EmailNotifications: true,
			PushNotifications:  true,
			MedicationAlerts:   true,
		},
	}

	_, err = database.Client.Collection("users").Doc(user.ID).Set(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"fullName": user.FullName,
		},
	})
}

func GoogleAuth(c *gin.Context) {
	var req models.GoogleAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// Verify Firebase ID token
	authClient, err := database.FirebaseApp.Auth(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize auth client"})
		return
	}

	token, err := authClient.VerifyIDToken(ctx, req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google token"})
		return
	}

	// Get user info from token
	userRecord, err := authClient.GetUser(ctx, token.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	// Check if user exists in Firestore
	iter := database.Client.Collection("users").Where("email", "==", userRecord.Email).Limit(1).Documents(ctx)
	defer iter.Stop()

	var user models.User

	doc, err := iter.Next()
	if err != nil && err != iterator.Done {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if err == iterator.Done {
		// Create new user
		user = models.User{
			ID:        userRecord.UID,
			Email:     userRecord.Email,
			FullName:  userRecord.DisplayName,
			Provider:  "google",
			GoogleID:  userRecord.UID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Settings: models.UserSettings{
				EmailNotifications: true,
				PushNotifications:  true,
				MedicationAlerts:   true,
			},
		}

		// Use custom name if provided
		if req.FullName != "" {
			user.FullName = req.FullName
		}

		_, err = database.Client.Collection("users").Doc(user.ID).Set(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}
	} else {
		// Update existing user - we might want to update some fields
		if err := doc.DataTo(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		// Update the user's information if needed (e.g., profile picture, display name)
		// For now, we'll just use the existing user data
		// You could add update logic here if needed
	}

	// Generate our JWT token
	jwtToken, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Google authentication has been done successfully",
		"token":   jwtToken,
		"user": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"fullName": user.FullName,
		},
	})
}
