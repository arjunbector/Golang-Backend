package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func DBInstance() *mongo.Client {
	// Load .env file if client is not initialized
	if client == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}

		mongoURI := os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			log.Fatal("MONGODB_URI environment variable is not set")
		}

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Create new client
		var err error
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
		if err != nil {
			log.Fatal("Error connecting to MongoDB:", err)
		}

	}
	fmt.Println("Connnected to mongoDb")
	return client
}

var Client *mongo.Client = DBInstance()