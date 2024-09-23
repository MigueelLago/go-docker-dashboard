package usecases

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// Struct represent DockerImage
type DockerImage struct {
	ID       string   `json:"id"`
	RepoTags []string `json:"repo_tags"`
	Size     int64    `json:"size"`
}

// Struct Docker Operations
type DockerUseCase struct {
	Client *client.Client
}

// Create new UseCase for Docker
func NewDockerUseCase(cli *client.Client) *DockerUseCase {
	return &DockerUseCase{Client: cli}
}

// List DockerImages
func (useCase *DockerUseCase) ListDockerImages() ([]DockerImage, error) {
	images, err := useCase.Client.ImageList(context.Background(), image.ListOptions{})

	if err != nil {
		return nil, fmt.Errorf("erro ao listar imagens: %w", err)
	}

	var dockerImages []DockerImage
	for _, img := range images {
		dockerImages = append(dockerImages, DockerImage{
			ID:       img.ID,
			RepoTags: img.RepoTags,
			Size:     img.Size,
		})
	}

	return dockerImages, nil
}
