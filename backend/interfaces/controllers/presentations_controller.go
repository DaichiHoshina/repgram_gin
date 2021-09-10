package controllers

import (
	"strconv"

	"github.com/DaichiHoshina/repgram_gin/backend/interfaces/database"
	"github.com/DaichiHoshina/repgram_gin/backend/usecase"
)

type PresentationsController struct {
	Interactor usecase.PresentationInteractor
}

func NewPresentationsController(db database.DB) *PresentationsController {
	return &PresentationsController{
		Interactor: usecase.PresentationInteractor{
			DB:           &database.DBRepository{DB: db},
			Presentation: &database.PresentationRepository{},
		},
	}
}

func (controller *PresentationsController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.GetPresentation(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, user)
}
