package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	songrequest "github.com/davecusatis/song-request-backend/song-request"
	"github.com/davecusatis/song-request-backend/song-request/api"
)

func initS3() *s3manager.Uploader {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	return s3manager.NewUploader(sess)
}

func main() {
	s3Uploader := initS3()
	api, err := api.NewAPI(s3Uploader)
	server, err := songrequest.NewServer(api)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

	server.Start()
}
