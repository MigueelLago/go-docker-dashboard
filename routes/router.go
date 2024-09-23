package routes

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func Initialize(dockerClient *client.Client) {

	// Define app Gin Engine
	app := gin.Default()

	// Init Routes
	initRoutes(app, dockerClient)

	// App listen
	app.Run(":8080")
}
