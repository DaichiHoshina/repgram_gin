package controllers

import (
	"strconv"

	"github.com/psychedelicnekopunch/gin-clean-architecture/app/interfaces/database"
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/usecase"
)

type UsersController struct {
    Interactor usecase.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
    return &UsersController{
        Interactor: usecase.UserInteractor{
            DB: &database.DBRepository{ DB: db },
            User: &database.UserRepository{},
        },
    }
}

func (controller *UsersController) Get(c Context) {

    id, _ := strconv.Atoi(c.Param("id"))

    user, res := controller.Interactor.Get(id)
    if res.Error != nil {
        c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
        return
    }
    c.JSON(res.StatusCode, NewH("success", user))
}
