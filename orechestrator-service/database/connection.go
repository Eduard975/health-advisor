package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Users    *mongo.Collection
	Activities *mongo.Collection
	HealthRecords *mongo.Collection
	ChatMessages *mongo.Collection
)

func ConnectDB() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017/health-harbor"
	}

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB!")

	Client = client
	db := client.Database("health_harbor")

	Users = db.Collection("users")
	Activities = db.Collection("activities")
	HealthRecords = db.Collection("health_records")
	ChatMessages = db.Collection("chat_messages")
}