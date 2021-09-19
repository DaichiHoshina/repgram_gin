package controllers

import (
	"strconv"

	"github.com/DaichiHoshina/repgram_gin/backend/interfaces/database"
	"github.com/DaichiHoshina/repgram_gin/backend/usecase"
	"github.com/dgrijalva/jwt-go"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (controller *UsersController) Show(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.UserById(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, user)
}

func (controller *UsersController) Login(c Context) {
	token, res := controller.Interactor.UserLogin(c)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, token)

}

type Claims struct {
	jwt.StandardClaims
}

func (controller *UsersController) Connect(c Context) {
	// cookie取得
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(400, "cookie is not found")
		return
	}

	// token取得
	token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.JSON(400, "user is not login")
		return
	}

	claims := token.Claims.(*Claims)
	// User IDを取得
	id := claims.Issuer
	id_int, _ := strconv.Atoi(id)

	user, res := controller.Interactor.UserById(id_int)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, user)
}
