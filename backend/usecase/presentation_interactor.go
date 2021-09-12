package usecase

import (
	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationInteractor struct {
	DB           DBRepository
	Presentation PresentationRepository
}

func (interactor *PresentationInteractor) Presentations() (presentation domain.Presentations, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	presentations, err := interactor.Presentation.FindAll(db)
	if err != nil {
		return domain.Presentations{}, NewResultStatus(404, err)
	}
	// var presentations_array []string
	// for _, s := range presentations {
	// 	presentations_array = append(presentations_array, s.BuildForGet())
	// }
	return presentations, NewResultStatus(200, nil)
}

func (interactor *PresentationInteractor) PresentationByID(id int) (presentation domain.PresentationForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundPresentation, err := interactor.Presentation.FindByID(db, id)
	if err != nil {
		return domain.PresentationForGet{}, NewResultStatus(404, err)
	}
	presentation = foundPresentation.BuildForGet()
	return presentation, NewResultStatus(200, nil)
}
