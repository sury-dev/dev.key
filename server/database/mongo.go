package database

import (
	"context"
	"log"
	"sync"

	"github.com/sury-dev/dev-key-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
	dbName              string
)

// This function is used to connect to the MongoDB database
func connect() {
	// Get the MongoDB URI and database name from the environment variables
	mongoURI := config.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName = config.GetEnv("MONGO_DB_NAME", "dev-key")

	// Create a new MongoDB client
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to the MongoDB database
	client, err := mongo.Connect(context.Background(), clientOptions)

	// If there is an error connecting to the MongoDB database, set the clientInstanceError and log the error
	if err != nil {
		clientInstanceError = err
		log.Fatal("Error connecting to MongoDB", err)
		return
	}

	// Set the client instance to the MongoDB client
	clientInstance = client

	log.Println("Connected to MongoDB")
}

// This function is used to get the MongoDB client
func GetMongoClient() (*mongo.Client, error) {
	// Connect to the MongoDB database if it is not already connected
	//Called only once
	mongoOnce.Do(connect)

	return clientInstance, clientInstanceError
}

// This function is used to ping the MongoDB database
// It is used to check if the MongoDB database is connected
func Ping() error {
	// Get the MongoDB client
	client, err := GetMongoClient()

	if err != nil {
		return err
	}

	// Ping the MongoDB database
	return client.Ping(context.Background(), nil)
}

// This function is used to get the MongoDB collection
func GetCollection(collectionName string) (*mongo.Collection, error) {
	// Get the MongoDB client
	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	// Get the MongoDB collection
	return client.Database(dbName).Collection(collectionName), nil
}
