package usecases

import "github.com/docker/docker/client"

type DockerUseCase struct {
	Client *client.Client
}

// Create new UseCase for Docker
func NewDockerUseCase(cli *client.Client) *DockerUseCase {
	return &DockerUseCase{Client: cli}
}
