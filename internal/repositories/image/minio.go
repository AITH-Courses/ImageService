package image

import (
	"context"
	"io"
	"path"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioImageRepository struct {
	client     *minio.Client
	bucketName string
}

func (repo *MinioImageRepository) AddOne(filename string, fileSize int64, reader io.Reader) (string, error) {
	contentType := "image/" + path.Ext(filename)
	_, putObjectError := repo.client.PutObject(context.Background(), repo.bucketName, filename, reader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if putObjectError != nil {
		return "", putObjectError
	}
	pathToFile := repo.bucketName + "/" + filename
	return pathToFile, nil
}

func createBucket(ctx context.Context, client *minio.Client, bucketName string) error {
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

func NewMinioImageRepository(ctx context.Context, host, port, user, password, bucketName string, useSSL bool) (*MinioImageRepository, error) {
	endpoint := host + ":" + port
	client, getMinioClientError := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: useSSL,
	})
	if getMinioClientError != nil {
		return nil, getMinioClientError
	}
	createBucketError := createBucket(ctx, client, bucketName)
	if createBucketError != nil {
		return nil, createBucketError
	}
	repo := &MinioImageRepository{
		client:     client,
		bucketName: bucketName,
	}
	return repo, nil
}
