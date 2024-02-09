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

	// Load HTML templates from the "templates" directory
	router.LoadHTMLGlob("templates/*")

	// User Routes
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.GET("/excel", controllers.GetUsersExcel)
		userGroup.GET("/report", controllers.GetUserHtml)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}
