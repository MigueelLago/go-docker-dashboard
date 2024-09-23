package routes

import "github.com/gin-gonic/gin"

func Initialize() {

	// Define app Gin Engine
	app := gin.Default()

	// Init Routes
	initRoutes(app)

	// App listen
	app.Run(":8080")
}
