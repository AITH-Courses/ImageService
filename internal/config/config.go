package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MinioHost         string
	MinioPort         string
	MinioBucketName   string
	MinioRootUser     string
	MinioRootPassword string
	MinioUseSSL       bool
}

func LoadConfig() (*Config, error) {
	loadConfigError := godotenv.Load()
	if loadConfigError != nil {
		return nil, loadConfigError
	}
	minioHost, envError := getEnv("MINIO_HOST")
	if envError != nil {
		return nil, envError
	}
	minioPort, envError := getEnv("MINIO_PORT")
	if envError != nil {
		return nil, envError
	}
	minioBucketName, envError := getEnv("MINIO_BUCKET_NAME")
	if envError != nil {
		return nil, envError
	}
	minioRootUser, envError := getEnv("MINIO_ROOT_USER")
	if envError != nil {
		return nil, envError
	}
	minioRootPassword, envError := getEnv("MINIO_ROOT_PASSWORD")
	if envError != nil {
		return nil, envError
	}
	minioUseSSL, envError := getEnvAsBool("MINIO_USE_SSL")
	if envError != nil {
		return nil, envError
	}
	config := &Config{
		MinioHost:         minioHost,
		MinioPort:         minioPort,
		MinioBucketName:   minioBucketName,
		MinioRootUser:     minioRootUser,
		MinioRootPassword: minioRootPassword,
		MinioUseSSL:       minioUseSSL,
	}
	return config, nil
}

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("not value for key=%s in env", key)
}

func getEnvAsBool(key string) (bool, error) {
	rawValue, keyError := getEnv(key)
	if keyError != nil {
		return false, keyError
	}
	value, parseBoolError := strconv.ParseBool(rawValue)
	if parseBoolError != nil {
		return false, parseBoolError
	}
	return value, nil
}
