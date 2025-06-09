package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	bucket := "nico-wedding"
	object := r.URL.Query().Get("object")

	// Initialize AWS S3 client but point it to GCS
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("auto"), // region is ignored but required
		Endpoint:         aws.String("https://storage.googleapis.com"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}))

	svc := s3.New(sess)

	// Set expiration time
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})

	// Generate pre-signed URL
	url, err := req.Presign(20 * time.Minute)
	if err != nil {
		log.Fatalf("Failed to sign request: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"signed_url":"%s"}`, url)
}
