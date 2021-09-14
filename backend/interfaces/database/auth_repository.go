package database

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type AuthRepository struct{}

func (repo *AuthRepository) FindByID(db *gorm.DB, id int) (user domain.User, err error) {
	user = domain.User{}
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.User{}, errors.New("user is not found")
	}
	return user, nil
}
