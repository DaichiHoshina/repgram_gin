package database

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type UserRepository struct{}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.User, err error) {
	user = domain.User{}
	db.First(&user, id)
	if err != nil {
		return domain.User{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(db *gorm.DB, email string) (user domain.User, err error) {
	user = domain.User{}

	db.Where("email = ?", email).First(&user)
	if err != nil {
		return domain.User{}, errors.New("user is not found")
	}
	log.Println(user)
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, postUser domain.User) (user domain.User, err error) {
	if result := db.Create(&postUser); result.Error != nil {
		return domain.User{}, errors.New("user is not found")
	}
	return postUser, nil
}
