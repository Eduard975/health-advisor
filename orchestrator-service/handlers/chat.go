package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"health-harbor-backend/database"
	"health-harbor-backend/models"
	"health-harbor-backend/utils"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func SendMessage(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// Save user message
	userMessage := models.ChatMessage{
		ID:        utils.GenerateID(),
		UserID:    userID,
		Text:      req.Message,
		Sender:    "user",
		Timestamp: time.Now(),
		SessionID: req.SessionID,
	}

	_, err := database.Client.Collection("chat_messages").Doc(userMessage.ID).Set(ctx, userMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save message"})
		return
	}

	// Generate AI response
	aiResponse := generateAIResponse(req.Message)

	// Save AI response
	aiMessage := models.ChatMessage{
		ID:        utils.GenerateID(),
		UserID:    userID,
		Text:      aiResponse,
		Sender:    "ai",
		Timestamp: time.Now(),
		SessionID: req.SessionID,
	}

	_, err = database.Client.Collection("chat_messages").Doc(aiMessage.ID).Set(ctx, aiMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save AI response"})
		return
	}

	// Return both messages
	response := models.ChatResponse{
		UserMessage: userMessage,
		AiMessage:   aiMessage,
	}

	c.JSON(http.StatusOK, response)
}

func GetChatHistory(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	// Optional: get session ID from query params
	sessionID := c.Query("sessionId")

	// Optional: limit and pagination
	limit, err := parseInt(c.Query("limit"), 50)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	ctx := context.Background()

	// Build query - only get messages for this specific user
	query := database.Client.Collection("chat_messages").Where("userId", "==", userID)
	if sessionID != "" {
		query = query.Where("sessionId", "==", sessionID)
	}

	// Find options - sort by timestamp descending (newest first) and limit
	iter := query.OrderBy("timestamp", firestore.Desc).Limit(limit).Documents(ctx)
	defer iter.Stop()

	var messages []models.ChatMessage
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch chat history"})
			return
		}

		var message models.ChatMessage
		if err := doc.DataTo(&message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode chat history"})
			return
		}
		messages = append(messages, message)
	}

	// Reverse to get chronological order (oldest first)
	messages = reverseMessages(messages)

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"count":    len(messages),
		"userId":   userID,
	})
}

func GetChatSessions(c *gin.Context) {
	userID := c.MustGet("userId").(string)

	ctx := context.Background()

	// Get all messages for this user to group by session
	iter := database.Client.Collection("chat_messages").
		Where("userId", "==", userID).
		OrderBy("timestamp", firestore.Desc).
		Documents(ctx)
	defer iter.Stop()

	// Group messages by session ID manually
	sessionMap := make(map[string]struct {
		LastMessage  models.ChatMessage
		MessageCount int
		FirstMessage time.Time
	})

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch chat sessions"})
			return
		}

		var message models.ChatMessage
		if err := doc.DataTo(&message); err != nil {
			continue // Skip invalid messages
		}

		sessionID := message.SessionID
		if sessionID == "" {
			sessionID = "default" // Handle messages without session ID
		}

		session, exists := sessionMap[sessionID]
		if !exists {
			session = struct {
				LastMessage  models.ChatMessage
				MessageCount int
				FirstMessage time.Time
			}{
				LastMessage:  message,
				MessageCount: 0,
				FirstMessage: message.Timestamp,
			}
		}

		session.MessageCount++
		if message.Timestamp.After(session.LastMessage.Timestamp) {
			session.LastMessage = message
		}
		if message.Timestamp.Before(session.FirstMessage) {
			session.FirstMessage = message.Timestamp
		}

		sessionMap[sessionID] = session
	}

	// Convert map to slice for response
	var sessions []map[string]interface{}
	for sessionID, sessionData := range sessionMap {
		session := map[string]interface{}{
			"sessionId":    sessionID,
			"lastMessage":  sessionData.LastMessage,
			"messageCount": sessionData.MessageCount,
			"firstMessage": sessionData.FirstMessage,
			"lastActivity": sessionData.LastMessage.Timestamp,
		}
		sessions = append(sessions, session)
	}

	// Sort sessions by last activity (newest first)
	// Simple bubble sort for small datasets - consider better sorting for large datasets
	for i := 0; i < len(sessions)-1; i++ {
		for j := 0; j < len(sessions)-i-1; j++ {
			timeJ := sessions[j]["lastActivity"].(time.Time)
			timeJ1 := sessions[j+1]["lastActivity"].(time.Time)
			if timeJ.Before(timeJ1) {
				sessions[j], sessions[j+1] = sessions[j+1], sessions[j]
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"sessions": sessions,
		"userId":   userID,
	})
}

// Helper function to generate AI responses
func generateAIResponse(userMessage string) string {
	// This is a simple mock - in production, integrate with actual AI service
	payload := map[string]string{"query": userMessage}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Sprintf("Error encoding JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:8000/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Sprintf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response: %v", err)
	}

	// Return response as string (you can unmarshal JSON if needed)
	return string(body)
}

// Helper function to reverse message order
func reverseMessages(messages []models.ChatMessage) []models.ChatMessage {
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages
}

// Helper function to parse integers with default value
func parseInt(s string, defaultValue int) (int, error) {
	if s == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(s)
}
