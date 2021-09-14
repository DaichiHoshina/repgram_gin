package controllers

import (
	"strconv"

	"github.com/DaichiHoshina/repgram_gin/backend/interfaces/database"
	"github.com/DaichiHoshina/repgram_gin/backend/usecase"
	"github.com/dgrijalva/jwt-go"
)

type AuthsController struct {
	Interactor usecase.AuthInteractor
}

func NewAuthsController(db database.DB) *AuthsController {
	return &AuthsController{
		Interactor: usecase.AuthInteractor{
			DB: &database.DBRepository{DB: db},
			// Auth: &database.AuthRepository{},
		},
	}
}

type Claims struct {
	jwt.StandardClaims
}

func (controller *AuthsController) Connect(c Context) {
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

	user, res := controller.Interactor.AuthConnect(id_int)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, user)
}
