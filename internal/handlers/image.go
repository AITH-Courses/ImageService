package handlers

import (
	"encoding/json"
	schemas "image_service/internal/schemas"
	services "image_service/internal/services"
	"log"
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
	log.Printf("Try to load file '%s' with size %d", header.Filename, header.Size)
	URL, serviceError := ih.imageService.SaveImage(path.Ext(header.Filename), header.Size, file)
	if serviceError != nil {
		if _, ok := serviceError.(*services.FileError); ok {
			log.Print("File error: " + serviceError.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			encodeError := json.NewEncoder(w).Encode(schemas.NewErrorResponse(serviceError.Error()))
			if encodeError != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		} else {
			log.Print("Service error: " + serviceError.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	log.Print("Image succesfully loaded: " + URL)
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
