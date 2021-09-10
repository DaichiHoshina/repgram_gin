package usecase

import (
	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationInteractor struct {
	DB   DBRepository
	Presentation PresentationRepository
}

func (interactor *PresentationInteractor) GetPresentation(id int) (user domain.PresentationForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Presentation の取得
	foundPresentation, err := interactor.Presentation.FindByID(db, id)
	if err != nil {
		return domain.PresentationForGet{}, NewResultStatus(404, err)
	}
	user = foundPresentation.GetPresentation()
	return user, NewResultStatus(200, nil)
}
