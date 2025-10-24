package models

import "time"

type Activity struct {
	ID          string    `firestore:"id" json:"id"`
	UserID      string    `firestore:"userId" json:"userId"`
	Type        string    `firestore:"type" json:"type"` // "steps", "heart_rate", "water", "sleep", "exercise"
	Value       float64   `firestore:"value" json:"value"`
	Unit        string    `firestore:"unit" json:"unit"`
	Description string    `firestore:"description" json:"description"`
	Date        time.Time `firestore:"date" json:"date"`
	CreatedAt   time.Time `firestore:"createdAt" json:"createdAt"`
}

type ActivitySummary struct {
	Steps          int     `json:"steps"`
	StepsGoal      int     `json:"stepsGoal"`
	HeartRate      int     `json:"heartRate"`
	Water          int     `json:"water"`
	WaterGoal      int     `json:"waterGoal"`
	Sleep          float64 `json:"sleep"`
	SleepGoal      float64 `json:"sleepGoal"`
	TotalSteps     int     `json:"totalSteps"`
	ActiveMinutes  int     `json:"activeMinutes"`
	CaloriesBurned int     `json:"caloriesBurned"`
}
