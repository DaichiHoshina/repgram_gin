package domain

import (
	"fmt"
	"os"
)

type Config struct {
	Aws struct {
		S3 struct {
			Region          string
			Bucket          string
			AccessKeyID     string
			SecretAccessKey string
		}
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.Aws.S3.Region = "ap-northeast-1"
	c.Aws.S3.Bucket = os.Getenv("BUCKET_NAME")
	c.Aws.S3.AccessKeyID = os.Getenv("AWS_ACCESS_KEY")
	c.Aws.S3.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	fmt.Println(c.Aws.S3.Bucket)
	fmt.Println(c.Aws.S3.AccessKeyID)
	fmt.Println(c.Aws.S3.SecretAccessKey)

	return c
}
