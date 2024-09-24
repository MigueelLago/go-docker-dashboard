package usecases

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

func (d *DockerUseCase) FindContainersByImage(imageName string) ([]DockerContainer, error) {

	containers, err := d.Client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var matchedContainers []DockerContainer
	for _, container := range containers {
		if container.Image == imageName {
			matchedContainer := DockerContainer{
				ID:      container.ID,
				Image:   container.Image,
				Command: container.Command,
				Created: container.Created,
				Names:   container.Names,
				Status:  container.Status,
			}
			matchedContainers = append(matchedContainers, matchedContainer)
		}
	}

	if len(matchedContainers) == 0 {
		return nil, fmt.Errorf("não existe container com esse nome de imagem: %s", imageName)
	}

	return matchedContainers, nil
}
