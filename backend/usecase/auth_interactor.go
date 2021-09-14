package usecase

import (
	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type AuthInteractor struct {
	DB   DBRepository
	Auth AuthRepository
}

func (interactor *AuthInteractor) AuthByID(id int) (auth domain.AuthForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundAuth, err := interactor.Auth.FindByID(db, id)
	if err != nil {
		return domain.AuthForGet{}, NewResultStatus(404, err)
	}
	auth = foundAuth.BuildForGet()
	return auth, NewResultStatus(200, nil)
}

func (interactor *AuthInteractor) Login(id int) (auth domain.AuthForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundAuth, err := interactor.Auth.FindByID(db, id)
	if err != nil {
		return domain.AuthForGet{}, NewResultStatus(404, err)
	}
	auth = foundAuth.BuildForGet()
	return auth, NewResultStatus(200, nil)
}

func (interactor *AuthInteractor) AuthConnect(id int) (auth domain.AuthForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundAuth, err := interactor.Auth.FindByID(db, id)
	if err != nil {
		return domain.AuthForGet{}, NewResultStatus(404, err)
	}
	auth = foundAuth.BuildForGet()
	return auth, NewResultStatus(200, nil)
}
