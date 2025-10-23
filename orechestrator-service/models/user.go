package models

import "time"

type User struct {
	ID           string       `firestore:"id" json:"id"`
	Email        string       `firestore:"email" json:"email"`
	Password     string       `firestore:"password,omitempty" json:"-"`
	FullName     string       `firestore:"fullName" json:"fullName"`
	DateOfBirth  time.Time    `firestore:"dateOfBirth" json:"dateOfBirth"`
	Gender       string       `firestore:"gender" json:"gender"`
	Height       float64      `firestore:"height" json:"height"`
	Weight       float64      `firestore:"weight" json:"weight"`
	BloodType    string       `firestore:"bloodType" json:"bloodType"`
	Allergies    string       `firestore:"allergies" json:"allergies"`
	Medications  string       `firestore:"medications" json:"medications"`
	Conditions   string       `firestore:"conditions" json:"conditions"`
	ProfileImage string       `firestore:"profileImage" json:"profileImage"`
	Settings     UserSettings `firestore:"settings" json:"settings"`
	CreatedAt    time.Time    `firestore:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time    `firestore:"updatedAt" json:"updatedAt"`
	Provider     string       `firestore:"provider" json:"provider"` // "email" or "google"
	GoogleID     string       `firestore:"googleId,omitempty" json:"-"`
}

type UserSettings struct {
	EmailNotifications bool `firestore:"emailNotifications" json:"emailNotifications"`
	PushNotifications  bool `firestore:"pushNotifications" json:"pushNotifications"`
	ActivityReminders  bool `firestore:"activityReminders" json:"activityReminders"`
	MedicationAlerts   bool `firestore:"medicationAlerts" json:"medicationAlerts"`
	TwoFactorAuth      bool `firestore:"twoFactorAuth" json:"twoFactorAuth"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullName" binding:"required"`
}

type GoogleAuthRequest struct {
	Token    string `json:"token" binding:"required"`
	FullName string `json:"fullName,omitempty"`
}
