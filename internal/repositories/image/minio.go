package image

import (
	"context"
	"io"
	"path"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioImageRepository struct {
	imageEndpointPrefix string
	client              *minio.Client
	bucketName          string
}

func (repo *MinioImageRepository) AddOne(filename string, fileSize int64, reader io.Reader) (string, error) {
	contentType := "image/" + path.Ext(filename)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // таймаут 1 секунда для добавления файла
	defer cancel()
	_, putObjectError := repo.client.PutObject(ctx, repo.bucketName, filename, reader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if putObjectError != nil {
		return "", putObjectError
	}
	pathToFile := repo.imageEndpointPrefix + "/" + repo.bucketName + "/" + filename
	return pathToFile, nil
}

func createBucket(client *minio.Client, bucketName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // таймаут 3 секунды на все действия
	defer cancel()
	bucketExists, bucketExistsError := client.BucketExists(ctx, bucketName)
	if bucketExistsError != nil {
		return bucketExistsError
	}
	if !bucketExists {
		makeBucketError := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if makeBucketError != nil {
			return makeBucketError
		}
	}
	policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Sid": "PublicRead",
				"Effect": "Allow",
				"Principal": "*",
				"Action": "s3:GetObject",
				"Resource": "arn:aws:s3:::` + bucketName + `/*"
			}
		]
	}`

	setPolicyError := client.SetBucketPolicy(ctx, bucketName, policy)
	if setPolicyError != nil {
		return setPolicyError
	}
	return nil
}

func NewMinioImageRepository(host, port, user, password, bucketName string, useSSL bool, imageEndpointPrefix string) (*MinioImageRepository, error) {
	endpoint := host + ":" + port
	client, getMinioClientError := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: useSSL,
	})
	if getMinioClientError != nil {
		return nil, getMinioClientError
	}
	createBucketError := createBucket(client, bucketName)
	if createBucketError != nil {
		return nil, createBucketError
	}
	repo := &MinioImageRepository{
		imageEndpointPrefix: imageEndpointPrefix,
		client:              client,
		bucketName:          bucketName,
	}
	return repo, nil
}
