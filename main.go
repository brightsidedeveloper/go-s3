package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	endpoint := os.Getenv("ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}

	fmt.Println("MinIO client initialized successfully.")

	bucketName := "test"
	filePath := "vim.txt"
	objectName := "vim.txt"

	info, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	fmt.Printf("File uploaded successfully: %s (%d bytes)\n", info.Key, info.Size)
}
