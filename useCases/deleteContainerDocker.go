package usecases

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

func (d *DockerUseCase) DeleteContainerDocker(containeID string) (string, error) {

	if err := d.Client.ContainerRemove(context.Background(), containeID, container.RemoveOptions{
		Force: true,
	}); err != nil {
		return "", fmt.Errorf("falha ao deletar o container %s: %w", containeID, err)
	}

	return "Container deletado com sucesso!", nil
}
