package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

func main() {
	fileTransfer := FileTransferToS3{
		AccessKeyId:     "ACCESS_KEY_ID",
		SecretAccessKey: "SECRET_ACCESS_KEY",
	}

	fileTransfer.PutToS3("./", "sample.jpg")
}

type FileTransferToS3 struct {
	AccessKeyId     string
	SecretAccessKey string
}

func (f *FileTransferToS3) PutToS3(path string, filename string) {
	file, err := os.Open(fmt.Sprintf("%s%s", path, filename))
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	cli := s3.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(f.AccessKeyId, f.SecretAccessKey, ""),
		Region:      "us-west-2",
	})

	bucketName := "bucket-name"
	resp, err := cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(awsutil.StringValue(resp))
}
