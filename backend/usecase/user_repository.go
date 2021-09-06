package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
}
