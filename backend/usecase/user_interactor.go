package usecase

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

func (interactor *UserInteractor) UserById(id int) (user domain.UserForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// User の取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UserForGet{}, NewResultStatus(404, err)
	}
	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) UserLogin(c Context) (token string, resultStatus *ResultStatus) {
	post := new(domain.User)
	if err := c.Bind(post); err != nil {
		c.JSON(400, "post error")
		return
	}

	db := interactor.DB.Connect()

	var user domain.User

	// User の取得
	user, err := interactor.User.FindByEmail(db, post.Email)
	if err != nil {
		c.JSON(400, "メールアドレスが存在しません")
		return
	}

	if user.ID == 0 {
		log.Print("データが存在しません")
		return "", NewResultStatus(404, nil)
	}

	// パスワードのチェック
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(post.Password)); err != nil {
		log.Print("パスワードが一致しません")
		return "", NewResultStatus(404, nil)
	}

	// JWTトークンを取得
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte("secret"))
	if err != nil {
		log.Print("トークンの取得に失敗しました")
		return "", NewResultStatus(404, nil)
	}

	// Cookieをセット
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = true
	cookie.HttpOnly = true
	c.SetCookie("user", cookie.Value, 3600, "/", "localhost", false, true)

	fmt.Println(c.Cookie("jwt"))

	return token, NewResultStatus(200, nil)
}
