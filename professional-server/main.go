package main

import (
	"context"
	"log"
	"user/app"
	"user/controllers"
	"user/db"
)

func main() {
	client, err := db.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Initialize MongoDB collection
	controllers.InitMongoDBCollection(client)

	// Create a new application
	application := app.NewApplication()

	// Set up routes
	application.SetupRouter()

	// Start the server
	application.Start()

}
