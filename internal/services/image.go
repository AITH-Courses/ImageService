package services

import (
	"errors"
	repository "image_service/internal/repositories/image"
	"io"

	"github.com/google/uuid"
)

type FileError struct {
	error
}

type ImageService struct {
	repo repository.ImageRepository
}

func (service *ImageService) SaveImage(extension string, fileSize int64, file io.Reader) (string, error) {
	extensionError := checkFileExtension(extension)
	if extensionError != nil {
		return "", extensionError
	}
	fileSizeError := checkFileSize(fileSize)
	if fileSizeError != nil {
		return "", fileSizeError
	}
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

func checkFileExtension(extension string) error {
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	for _, allowedExtension := range allowedExtensions {
		if extension == allowedExtension {
			return nil
		}
	}
	return &FileError{error: errors.New("недопустимое расширение файла")}
}

func checkFileSize(fileSize int64) error {
	// 4 Mb
	maxSize := 4 * 1024 * 1024
	if fileSize > int64(maxSize) {
		return &FileError{error: errors.New("размер изображения не должен превышать 4 Мб")}
	}
	return nil
}

func NewImageService(repo repository.ImageRepository) ImageService {
	return ImageService{
		repo: repo,
	}
}
