package usecases

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Struct Docker Operations
type DockerUseCase struct {
	Client *client.Client
}

// Struct represent DockerImage
type DockerContainer struct {
	ID      string   `json:"id"`
	Image   string   `json:"image"`
	Command string   `json:"command"`
	Created int64    `json:"created"`
	Names   []string `json:"names"`
	Status  string   `json:"status"`
}

// Create new UseCase for Docker
func NewDockerUseCase(cli *client.Client) *DockerUseCase {
	return &DockerUseCase{Client: cli}
}

// List DockerImages
func (useCase *DockerUseCase) ListDockerContainers(onlyRunning bool) ([]DockerContainer, error) {

	containers, err := useCase.Client.ContainerList(context.Background(), container.ListOptions{
		All: !onlyRunning,
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao listar imagens: %w", err)
	}

	var dockerContainers []DockerContainer
	for _, container := range containers {
		dockerContainer := DockerContainer{
			ID:      container.ID,
			Image:   container.Image,
			Command: container.Command,
			Created: container.Created,
			Names:   container.Names,
			Status:  container.Status,
		}
		dockerContainers = append(dockerContainers, dockerContainer)
	}

	return dockerContainers, nil
}
