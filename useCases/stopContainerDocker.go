package usecases

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

func (d *DockerUseCase) StopContainerDocker(containeID string) (string, error) {

	if err := d.Client.ContainerStop(context.Background(), containeID, container.StopOptions{}); err != nil {
		return "", fmt.Errorf("falha ao parar container %s: %w", containeID, err)
	}

	return "Container parado com sucesso!", nil
}
