package main

import (
	"github.com/DaichiHoshina/repgram_gin/backend/infrastructure"
)

func main() {
	db := infrastructure.NewDB()

	r := infrastructure.NewRouting(db)

	r.Run()
}
