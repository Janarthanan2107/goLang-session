package app

import (
	"fmt"
	"user/routes"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Router *gin.Engine
}

func NewApplication() *Application {
	return &Application{
		Router: routes.SetupUserRouter(),
	}
}

func (app *Application) Start() {
	// Print a message indicating that the server is starting
	fmt.Println("=============================")
	fmt.Println("Server is starting...")
	fmt.Println("=============================")

	// Start the server
	fmt.Println("Server is started on the port http://localhost:8080")
	fmt.Println("=============================")
	app.Router.Run(":8080")
}
