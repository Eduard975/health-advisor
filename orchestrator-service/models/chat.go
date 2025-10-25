package models

import "time"

type ChatMessage struct {
	ID        string    `firestore:"id" json:"id"`
	UserID    string    `firestore:"userId" json:"userId"`
	Text      string    `firestore:"text" json:"text"`
	Sender    string    `firestore:"sender" json:"sender"` // "user" or "ai"
	Timestamp time.Time `firestore:"timestamp" json:"timestamp"`
	SessionID string    `firestore:"sessionId,omitempty" json:"sessionId"` // Optional: for grouping chats into sessions
}

type ChatHistoryItem struct {
	Sender    string    `json:"sender"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type ChatRequest struct {
	Message   string            `json:"message"`
	SessionID string            `json:"sessionId"`
	History   []ChatHistoryItem `json:"history"`
}

type ChatResponse struct {
	UserMessage ChatMessage `json:"userMessage"`
	AiMessage   ChatMessage `json:"aiMessage"`
}
