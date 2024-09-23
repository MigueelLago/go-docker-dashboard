package usecases

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

func (d *DockerUseCase) StartContainerDocker(containeID string) (string, error) {

	if err := d.Client.ContainerStart(context.Background(), containeID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("falha ao iniciar container %s: %w", containeID, err)
	}

	return "Container iniciado com sucesso!", nil
}
