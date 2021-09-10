package usecase

import (
	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

func (interactor *UserInteractor) GetUser(id int) (user domain.UserForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// User の取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UserForGet{}, NewResultStatus(404, err)
	}
	user = foundUser.GetUser()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) GetUsers(id int) (user domain.UserForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// User の取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UserForGet{}, NewResultStatus(404, err)
	}
	user = foundUser.GetUser()
	return user, NewResultStatus(200, nil)
}
