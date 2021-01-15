package Routes

import (
	"github.com/gin-gonic/gin"
	"mygithub.com/Logic"
)

func HandleCRUD() {
	// Use the default router in Gin
	router := gin.Default()

	///////////////////////////////////////////////////////////////
	// Below are all of the endpoints we'll need for Chatty-GoGo //
	///////////////////////////////////////////////////////////////

	api := router.Group("/v1")  // add versioning
	{
		// Home page
		api.GET("/welcome", Logic.HomeLink)
		// Register a new user
		api.POST("/add/user", Logic.UserRegistration)
		// List all registered users
		api.GET("/list/users", Logic.ListUsers)
		// Remove a registered user
		api.DELETE("/remove/:username", Logic.RemoveUser)
		// Update a user's data with a subject, message, and messagefrom
		api.PUT("/update/:username", Logic.UpdateMessageforUser)
	}

	// Start and run
	router.Run(":3001")  
}