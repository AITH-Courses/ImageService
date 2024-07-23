package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	repositories "image_service/internal/repositories/image"
	services "image_service/internal/services"

	config "image_service/internal/config"
	handlers "image_service/internal/handlers"
)

func main() {
	config, loadConfigError := config.LoadConfig()
	if loadConfigError != nil {
		fmt.Printf("%v", loadConfigError)
		return
	}
	imageRepo, createRepoError := repositories.NewMinioImageRepository(context.Background(), config.MinioHost, config.MinioPort, config.MinioRootUser, config.MinioRootPassword, config.MinioBucketName, config.MinioUseSSL)
	if createRepoError != nil {
		fmt.Printf("%v", createRepoError)
		return
	}
	imageService := services.NewImageService(imageRepo)
	imageHandler := handlers.NewImageHandler(imageService)
	router := mux.NewRouter()
	router.HandleFunc("/image", imageHandler.AddImage).
		Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
