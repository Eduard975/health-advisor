package database

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	FirebaseApp *firebase.App
	Client      *firestore.Client
)

func InitFirebase() {
	ctx := context.Background()

	// Use service account key file from environment variable
	serviceAccountKey := os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY")
	if serviceAccountKey == "" {
		log.Fatal("FIREBASE_SERVICE_ACCOUNT_KEY environment variable is required")
	}

	config := &firebase.Config{
		ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
	}

	opt := option.WithCredentialsJSON([]byte(serviceAccountKey))
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FirebaseApp = app

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing Firestore client: %v\n", err)
	}

	Client = client
	log.Println("Connected to Firebase Firestore!")
}

func CloseFirebase() {
	if Client != nil {
		Client.Close()
	}
}
