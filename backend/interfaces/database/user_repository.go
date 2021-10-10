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
		log.Println("ユーザーが見つかりませんでした。")
		return domain.User{}, errors.New("ユーザーが見つかりませんでした。")
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(db *gorm.DB, email string) (user domain.User, err error) {
	user = domain.User{}

	db.Where("email = ?", email).First(&user)
	if err != nil {
		log.Println("ユーザーが見つかりませんでした。")
		return domain.User{}, errors.New("ユーザーが作成出来ませんでした。")
	}
	log.Println(user)
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, postUser domain.User) (user domain.User, err error) {
	if result := db.Create(&postUser); result.Error != nil {
		log.Println("ユーザーが作成出来ませんでした。")
		return domain.User{}, errors.New("ユーザーが作成出来ませんでした。")
	}
	return postUser, nil
}
