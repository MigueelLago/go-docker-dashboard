package routes

import (
	"github.com/MigueelLago/go-docker-dashboard/handlers"
	usecases "github.com/MigueelLago/go-docker-dashboard/useCases"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, dockerClient *client.Client) {

	basePath := "/api"

	// Init UseCase
	dockerUseCase := usecases.NewDockerUseCase(dockerClient)

	// Init Handler
	dockerHandler := handlers.NewDockerHandler(dockerUseCase)

	docker := router.Group(basePath)
	{
		docker.GET("/containers", dockerHandler.ListContainers)
		docker.DELETE("/container/:containerID", dockerHandler.DeleteContainerDocker)
		docker.POST("/container/:containerID", dockerHandler.StartContainerDocker)
		docker.POST("/container/stop/:containerID", dockerHandler.StopContainerDocker)
	}
}
