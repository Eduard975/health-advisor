package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

func GenerateID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("Error generating ID: %v", err)
		// Fallback to timestamp-based ID
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return fmt.Sprintf("%x", b)
}
