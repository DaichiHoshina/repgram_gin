package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.User, err error)
	FindByEmail(db *gorm.DB, email string) (user domain.User, err error)
}
