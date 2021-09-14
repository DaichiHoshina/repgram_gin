package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type AuthRepository interface {
	FindByID(db *gorm.DB, id int) (auth domain.Auth, err error)
}
