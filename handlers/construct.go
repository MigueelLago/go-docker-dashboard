package handlers

import usecases "github.com/MigueelLago/go-docker-dashboard/useCases"

type DockerHandler struct {
	UseCase *usecases.DockerUseCase
}

// Create instance
func NewDockerHandler(useCase *usecases.DockerUseCase) *DockerHandler {
	return &DockerHandler{
		UseCase: useCase,
	}
}
