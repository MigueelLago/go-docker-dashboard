package handlers

import (
	"net/http"
	"strconv"

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

	onlyRunningQuery := ctx.Query("only_running")
	onlyRunning, err := strconv.ParseBool(onlyRunningQuery)
	if err != nil {
		onlyRunning = false
	}

	dockerContainers, err := h.UseCase.ListDockerContainers(onlyRunning)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":        "Failed to list Docker Images",
			"statusCode": 500,
		})
		return
	}

	if onlyRunning && len(dockerContainers) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":        "Não existem containers rodando atualmente",
			"statusCode": 200,
		})
		return
	}

	ctx.JSON(http.StatusOK, dockerContainers)
}
