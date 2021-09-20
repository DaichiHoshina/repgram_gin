package main

import (
	"github.com/DaichiHoshina/repgram_gin/backend/infrastructure"
)

func main() {
	db := infrastructure.NewDB()

	awsS3 := infrastructure.NewAwsS3()

	r := infrastructure.NewRouting(db, awsS3)

	r.Run()
}
