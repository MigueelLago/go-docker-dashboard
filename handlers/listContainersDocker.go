package handlers

import (
	"net/http"

	usecases "github.com/MigueelLago/go-docker-dashboard/useCases"
	"github.com/gin-gonic/gin"
)

type DockerHandler struct {
	UseCase *usecases.DockerUseCase
}

// Create instance
func NewDockerHandler(useCase *usecases.DockerUseCase) *DockerHandler {
	return &DockerHandler{
		UseCase: useCase,
	}
}

// List Images and return JSON format
func (h *DockerHandler) ListContainers(ctx *gin.Context) {
	dockerImages, err := h.UseCase.ListDockerImages()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":        "Failed to list Docker Images",
			"statusCode": 500,
		})
	}

	ctx.JSON(http.StatusOK, dockerImages)
}
