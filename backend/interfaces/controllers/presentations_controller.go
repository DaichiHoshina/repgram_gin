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

func (controller *PresentationsController) Index(c Context) {
	presentation, res := controller.Interactor.Presentations()
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, presentation)
}

func (controller *PresentationsController) Show(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	presentation, res := controller.Interactor.PresentationByID(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, presentation)
}

func (controller *PresentationsController) Create(c Context, awsS3 interface{}) {
	presentation, res := controller.Interactor.PresentationCreate(c, awsS3)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, presentation)
}
