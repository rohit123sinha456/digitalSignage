package objectstore

import (
	"context"
	"fmt"
	"log"
	"errors"
	"mime/multipart"
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

func createOrGetBucketIfExists(ctx context.Context, minioClient *minio.Client, bucketName string) (bool,error) {
	policy := `{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": "*",
                "Action": "s3:GetObject",
                "Resource": "arn:aws:s3:::` + bucketName + `/*"
            }
        ]
    }`
	doesExists,err := minioClient.BucketExists(ctx,bucketName)
	if err != nil{
		return false,err
	}
	
	if doesExists == true{
		return true,nil
	}

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		return false,err
	}
    // Set the bucket policy
    err = minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
    if err != nil {
       return false,err
    }
	fmt.Println("Successfully created UserBucket.")
	return true,nil
}


func CreateBucket(ctx context.Context, minioClient *minio.Client, bucketName string) error {
	policy := `{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": "*",
                "Action": "s3:GetObject",
                "Resource": "arn:aws:s3:::` + bucketName + `/*"
            }
        ]
    }`
	doesExists,err := minioClient.BucketExists(ctx,bucketName)
	if err != nil{
		return err
	}
	
	if doesExists == true{
		return nil
	}
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		return err
	}
	err = minioClient.SetBucketPolicy(ctx, bucketName, policy)
    if err != nil {
       return err
    }
	fmt.Println("Successfully created UserBucket.")
	return nil
}

func DeleteBucket(ctx context.Context, minioClient *minio.Client, bucketName string) error {
	err := minioClient.RemoveBucket(ctx, bucketName)
	if err != nil {
		return err
	}
	fmt.Println("Successfully Deleted UserBucket.")
	return nil
}

func StoreFile(ctx context.Context, minioClient *minio.Client, userBucketname string, filedata *multipart.FileHeader) error {
	
	bucketexists,err := createOrGetBucketIfExists(ctx,minioClient,userBucketname)
	if err != nil {
		return err
	}
	// Get Buffer from file
	log.Printf("Opening File")
    buffer, err := filedata.Open()
    if err != nil {
        return err
    }
    defer buffer.Close()

	objectName := filedata.Filename
    // fileBuffer := buffer
    contentType := filedata.Header["Content-Type"][0]
    fileSize := filedata.Size

	log.Printf("Uploading File with PutObject")
	log.Printf(userBucketname)
	log.Printf(objectName)
	log.Printf(contentType)
	log.Printf("%d",fileSize)

	if minioClient == nil {
		log.Printf("Minio Clinet is nil")
	}
	if buffer == nil {
		log.Printf("File is nil")
	}
    // Upload the zip file with PutObject
	if bucketexists == true {
		info, err := minioClient.PutObject(ctx, userBucketname, objectName, buffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			return err
		}
		log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
		return nil
	} else {
		return errors.New("Bucket Doesn't Exists and can't create bucket")
	}
}