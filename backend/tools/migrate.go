package main

import (
	"os"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
	"github.com/DaichiHoshina/repgram_gin/backend/infrastructure"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// マイグレーション
func main() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			logrus.Fatalf("Error loading env: %v", err)
		}
	}

	db := infrastructure.NewDB()

	db.Connection.AutoMigrate(&domain.Users{})
	// db.AutoMigrate(&model.Presentation{})
	// db.AutoMigrate(&model.Like{})
}
