package usecases

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

// Struct represent DockerImage
type DockerContainer struct {
	ID      string   `json:"id"`
	Image   string   `json:"image"`
	Command string   `json:"command"`
	Created int64    `json:"created"`
	Names   []string `json:"names"`
	Status  string   `json:"status"`
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
