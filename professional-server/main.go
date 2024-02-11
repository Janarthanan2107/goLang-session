package main

import (
	"context"
	"log"
	"user/app"
	"user/controllers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := ConnectToMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Initialize MongoDB collection
	controllers.InitMongoDBCollection(client)

	// Create a new application
	application := app.NewApplication()

	// Start the server
	application.Start()
}

// ConnectToMongoDB connects to MongoDB and returns a client instance
func ConnectToMongoDB() (*mongo.Client, error) {
	// MongoDB connection URI
	uri := "mongodb+srv://janarthanan:janarthanan2103@cluster-db.sndm3lz.mongodb.net/practice?retryWrites=true&w=majority"

	// Create a new MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check if the connection was successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
