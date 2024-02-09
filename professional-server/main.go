package main

import "user/app"

func main() {
	// Create a new application
	application := app.NewApplication()

	// Start the server
	application.Start()
}
