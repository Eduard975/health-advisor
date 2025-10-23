package handlers

import (
	"context"
	"net/http"
	"time"

	"health-harbor-backend/database"
	"health-harbor-backend/models"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	ctx := context.Background()
	doc, err := database.Client.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateProfile(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	var updateData struct {
		FullName    string    `json:"fullName"`
		DateOfBirth time.Time `json:"dateOfBirth"`
		Gender      string    `json:"gender"`
		Height      float64   `json:"height"`
		Weight      float64   `json:"weight"`
		BloodType   string    `json:"bloodType"`
		Allergies   string    `json:"allergies"`
		Medications string    `json:"medications"`
		Conditions  string    `json:"conditions"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	_, err := database.Client.Collection("users").Doc(userID).Update(ctx, []firestore.Update{
		{Path: "fullName", Value: updateData.FullName},
		{Path: "dateOfBirth", Value: updateData.DateOfBirth},
		{Path: "gender", Value: updateData.Gender},
		{Path: "height", Value: updateData.Height},
		{Path: "weight", Value: updateData.Weight},
		{Path: "bloodType", Value: updateData.BloodType},
		{Path: "allergies", Value: updateData.Allergies},
		{Path: "medications", Value: updateData.Medications},
		{Path: "conditions", Value: updateData.Conditions},
		{Path: "updatedAt", Value: time.Now()},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func UpdateSettings(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	var settings models.UserSettings
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	_, err := database.Client.Collection("users").Doc(userID).Update(ctx, []firestore.Update{
		{Path: "settings", Value: settings},
		{Path: "updatedAt", Value: time.Now()},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
}
