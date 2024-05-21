package objectstore

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rohit123sinha456/digitalSignage/config"

)

// var minioClient *minio.Client

func ConnectObjectStore() *minio.Client {
	var endpoint string = config.GetEnvbyKey("APPOSURL")                 //"localhost:9000"
	var accessKeyID string = config.GetEnvbyKey("APPOSACCKEY")           //"xiGpoR8ggd6gd3c47v0C"
	var secretAccessKey string = config.GetEnvbyKey("APPOSSECRETACCKEY") //"d1LtofnWlJflsXzKSu3h01o5WvIwZcnkqET7QyTd"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("minio Connected") // minioClient is now set up
	return minioClient
}

func CreateBucket(ctx context.Context, minioClient *minio.Client, bucketName string) error {
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		return err
	}
	fmt.Println("Successfully created UserBucket.")
	return nil
}
