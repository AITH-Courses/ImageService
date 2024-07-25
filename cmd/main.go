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
	web "image_service/internal/web"
)

func main() {
	config, loadConfigError := config.LoadConfig()
	if loadConfigError != nil {
		fmt.Printf("%v", loadConfigError)
		return
	}
	imageRepo, createRepoError := repositories.NewMinioImageRepository(context.Background(), config.MinioHost, config.MinioPort, config.MinioRootUser, config.MinioRootPassword, config.MinioBucketName, config.MinioUseSSL, config.ImageEndpointPrefix)
	if createRepoError != nil {
		fmt.Printf("%v", createRepoError)
		return
	}
	imageService := services.NewImageService(imageRepo)
	imageHandler := handlers.NewImageHandler(imageService)
	router := mux.NewRouter()
	router.HandleFunc("/admin/images", imageHandler.AddImage).
		Methods("POST")

	cors := web.NewCORS(config.AllowedOrigins)
	router.Use(cors.Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
