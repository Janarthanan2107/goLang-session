package app

import (
	"fmt"
	"net/http"
	"user/routes"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Router *gin.Engine
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) SetupRouter() {
	// Create a main router
	router := gin.Default()

	// Set up root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the server",
		})
	})

	// Set up user routes
	userRouter := routes.SetupUserRouter()
	router.Any("/users/*any", gin.WrapH(userRouter))

	// example:
	// Set up product routes
	// productRouter := routes.SetupProductRouter()
	// router.Any("/products/*any", gin.WrapH(productRouter))

	app.Router = router
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
