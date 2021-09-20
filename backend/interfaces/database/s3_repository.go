package database

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct{}

func (a s3manager) UploadTest(file multipart.File, fileName string, extension string) (url string, err error) {

	if fileName == "" {
		return "", errors.New("fileName is required")
	}

	var contentType string

	switch extension {
	case "jpg":
		contentType = "image/jpeg"
	case "jpeg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	default:
		return "", errors.New("this extension is invalid")
	}

	// Upload the file to S3.
	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		// ACL の設定は重要
		ACL:    aws.String("public-read"),
		Body:   file,
		Bucket: aws.String(a.Config.Aws.S3.Bucket),
		// content-type の設定も忘れずに
		ContentType: aws.String(contentType),
		Key:         aws.String(a.Keys.Test + "/" + fileName + "." + extension),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}
