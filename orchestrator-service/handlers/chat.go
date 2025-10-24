package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
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
	responses := map[string]string{
		"hello":      "Hello! I'm Health Harbor AI, your personal health advisor. How can I assist you with your health concerns today?",
		"headache":   "I understand you're experiencing a headache. This could be due to various factors like stress, dehydration, or tension. Make sure to drink plenty of water, rest in a quiet room, and consider over-the-counter pain relief if appropriate. If the headache is severe, persistent, or accompanied by other symptoms like vision changes or fever, please consult a healthcare professional immediately.",
		"sleep":      "Sleep is crucial for overall health. Adults typically need 7-9 hours of quality sleep per night. Maintain a consistent sleep schedule, create a relaxing bedtime routine, and avoid screens before bed. If you're having persistent sleep issues, it's best to consult with a healthcare provider.",
		"fever":      "A fever is often a sign that your body is fighting an infection. Make sure to stay hydrated, rest, and monitor your temperature. If the fever is high (above 103°F/39.4°C), lasts more than 3 days, or is accompanied by severe symptoms like difficulty breathing or a stiff neck, seek medical attention immediately.",
		"cough":      "A cough can be caused by various factors including colds, allergies, or respiratory infections. Stay hydrated, use a humidifier, and consider over-the-counter remedies for symptom relief. If your cough persists for more than 3 weeks, is accompanied by chest pain, or you're coughing up blood, please see a doctor.",
		"stress":     "Stress can have significant impacts on both mental and physical health. Practice relaxation techniques like deep breathing, meditation, or gentle exercise. Ensure you're getting enough sleep and maintaining a balanced diet. If stress is affecting your daily life, consider speaking with a mental health professional.",
		"diet":       "A balanced diet is essential for good health. Focus on whole foods like fruits, vegetables, lean proteins, and whole grains. Stay hydrated and limit processed foods and added sugars. For personalized nutrition advice, consider consulting with a registered dietitian.",
		"exercise":   "Regular physical activity is important for maintaining health. Aim for at least 150 minutes of moderate exercise per week. Start slowly if you're new to exercise and choose activities you enjoy. Always consult with a healthcare provider before starting a new exercise program, especially if you have existing health conditions.",
		"calories":   "Multiple exercisies are available for this problem",
		"Stationary": "For this type of problem It is recommended to have a strict died for a deficitary calories intake",
	}

	// Simple keyword matching - replace with actual AI integration
	lowerMsg := strings.ToLower(userMessage)
	for keyword, response := range responses {
		if strings.Contains(lowerMsg, keyword) {
			return response
		}
	}

	// Default response
	return "Thank you for sharing your health concern. I'm analyzing this and will provide general guidance. Remember, I can offer health information but for specific medical advice, diagnosis, or treatment, please consult with a qualified healthcare professional. Could you tell me more about your symptoms or concerns?"
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
