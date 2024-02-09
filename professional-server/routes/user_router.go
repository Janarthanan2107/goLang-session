package routes

import (
	"net/http"
	"user/controllers"

	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome to the server")
}

func SetupUserRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", root)

	// User Routes
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}
