package services

import (
	repository "image_service/internal/repositories/image"
	"io"

	"github.com/google/uuid"
)

type ImageService struct {
	repo repository.ImageRepository
}

func (service *ImageService) SaveImage(extension string, fileSize int64, file io.Reader) (string, error) {
	generatedUUID, uuidError := uuid.NewRandom()
	if uuidError != nil {
		return "", uuidError
	}
	filename := generatedUUID.String() + extension
	pathToFile, repositoryError := service.repo.AddOne(filename, fileSize, file)
	if repositoryError != nil {
		return "", repositoryError
	}
	return pathToFile, nil
}

func NewImageService(repo repository.ImageRepository) ImageService {
	return ImageService{
		repo: repo,
	}
}
