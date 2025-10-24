package handlers

import (
	"context"
	"net/http"
	"time"

	"orchestrator-service/database"
	"orchestrator-service/models"
	"orchestrator-service/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GetHealthRecords(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	ctx := context.Background()
	iter := database.Client.Collection("health_records").Where("userId", "==", userID).Documents(ctx)

	var records []models.HealthRecord
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch health records"})
			return
		}

		var record models.HealthRecord
		if err := doc.DataTo(&record); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode health records"})
			return
		}
		records = append(records, record)
	}

	c.JSON(http.StatusOK, records)
}

func CreateHealthRecord(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	var record models.HealthRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = utils.GenerateID()
	record.UserID = userID
	record.CreatedAt = time.Now()

	ctx := context.Background()

	_, err := database.Client.Collection("health_records").Doc(record.ID).Set(ctx, record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create health record"})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func DeleteHealthRecord(c *gin.Context) {
	userID := c.MustGet("userId").(string)
	recordID := c.Param("id")

	ctx := context.Background()

	// Verify the record belongs to the user
	doc, err := database.Client.Collection("health_records").Doc(recordID).Get(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Health record not found"})
		return
	}

	var record models.HealthRecord
	if err := doc.DataTo(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode health record"})
		return
	}

	if record.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	_, err = database.Client.Collection("health_records").Doc(recordID).Delete(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete health record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Health record deleted successfully"})
}
