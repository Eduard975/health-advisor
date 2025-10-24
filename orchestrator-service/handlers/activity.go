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

func GetActivity(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	ctx := context.Background()

	// Get today's activities
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	iter := database.Client.Collection("activities").
		Where("userId", "==", userID).
		Where("date", ">=", today).
		Where("date", "<", tomorrow).
		Documents(ctx)

	var activities []models.Activity
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch activities"})
			return
		}

		var activity models.Activity
		if err := doc.DataTo(&activity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode activities"})
			return
		}
		activities = append(activities, activity)
	}

	c.JSON(http.StatusOK, activities)
}

func CreateActivity(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity.ID = utils.GenerateID()
	activity.UserID = userID
	activity.CreatedAt = time.Now()

	ctx := context.Background()

	_, err := database.Client.Collection("activities").Doc(activity.ID).Set(ctx, activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create activity"})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

func GetActivitySummary(c *gin.Context) {
	// Mock data - you can implement real aggregation later
	summary := models.ActivitySummary{
		Steps:          8432,
		StepsGoal:      10000,
		HeartRate:      72,
		Water:          6,
		WaterGoal:      8,
		Sleep:          7.5,
		SleepGoal:      8,
		TotalSteps:     67845,
		ActiveMinutes:  245,
		CaloriesBurned: 2840,
	}

	c.JSON(http.StatusOK, summary)
}
