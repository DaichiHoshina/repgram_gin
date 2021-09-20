package infrastructure

import (
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
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.DB.Production.Host = "repgram_db"
	c.DB.Production.Username = "user"
	c.DB.Production.Password = "password"
	c.DB.Production.DBName = "mydb"

	c.DB.Test.Host = "repgram_db"
	c.DB.Test.Username = "user"
	c.DB.Test.Password = "password"
	c.DB.Test.DBName = "mydb"

	c.Routing.Port = ":3001"

	c.Aws.S3.Region = "ap-northeast-1"
	c.Aws.S3.Bucket = os.Getenv("BUCKET_NAME")
	c.Aws.S3.AccessKeyID = os.Getenv("AWS_ACCESS_KEY")
	c.Aws.S3.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	return c
}
