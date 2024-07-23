package handlers

import (
	"encoding/json"
	schemas "image_service/internal/schemas"
	services "image_service/internal/services"
	"net/http"
	"path"
)

type ImageHandler struct {
	imageService services.ImageService
}

func (ih *ImageHandler) AddImage(w http.ResponseWriter, r *http.Request) {
	file, header, checkFileError := r.FormFile("file")
	if checkFileError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	URL, serviceError := ih.imageService.SaveImage(path.Ext(header.Filename), header.Size, file)
	if serviceError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encodeError := json.NewEncoder(w).Encode(schemas.NewImageAdded(URL))
	if encodeError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewImageHandler(imageService services.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}
