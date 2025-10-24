package models

import "time"

type HealthRecord struct {
	ID          string    `firestore:"id" json:"id"`
	UserID      string    `firestore:"userId" json:"userId"`
	Title       string    `firestore:"title" json:"title"`
	Date        time.Time `firestore:"date" json:"date"`
	Doctor      string    `firestore:"doctor" json:"doctor"`
	Type        string    `firestore:"type" json:"type"`     // "checkup", "lab_work", "specialist", "immunization"
	Status      string    `firestore:"status" json:"status"` // "completed", "scheduled", "cancelled"
	Description string    `firestore:"description" json:"description"`
	FileURL     string    `firestore:"fileUrl" json:"fileUrl"`
	CreatedAt   time.Time `firestore:"createdAt" json:"createdAt"`
}
