package usecase

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		log.Print("ユーザーが見つかりませんでした")
		return domain.UserForGet{}, NewResultStatus(400, err)
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

	user, err := interactor.User.FindByEmail(db, post.Email)
	if err != nil {
		c.JSON(400, "メールアドレスが存在しません")
		return
	}

	// パスワードのチェック
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(post.Password)); err != nil {
		log.Print("パスワードが一致しません")
		return "", NewResultStatus(400, nil)
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
		return "", NewResultStatus(400, nil)
	}

	// Cookieをセット
	cookie := new(http.Cookie)
	cookie.Value = token

	c.SetSameSite(http.SameSiteNoneMode)

	if os.Getenv("ENV") == "local" {
		log.Println("cookieをセットする")
		c.SetCookie("jwt", cookie.Value, 3600, "/", "localhost", true, true)
	}

	if os.Getenv("ENV") == "production" {
		log.Println("productionでcookieをセットする")
		c.SetCookie("jwt", cookie.Value, 3600, "/", "repgram-api.net", true, true)
	}

	return token, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) UserLogout(c Context) (resultStatus *ResultStatus) {
	// Cookieをセット
	cookie := new(http.Cookie)
	cookie.Value = ""

	c.SetSameSite(http.SameSiteNoneMode)

	if os.Getenv("ENV") == "local" {
		log.Println("cookieをセットする")
		c.SetCookie("jwt", cookie.Value, 3600, "/", "localhost", true, true)
	}

	if os.Getenv("ENV") == "production" {
		log.Println("productionでcookieをセットする")
		c.SetCookie("jwt", cookie.Value, 3600, "/", "repgram-api.net", true, true)
	}

	return NewResultStatus(200, nil)
}

func (interactor *UserInteractor) UserCreate(c Context) (user domain.User, resultStatus *ResultStatus) {
	post := new(domain.User)
	if err := c.Bind(post); err != nil {
		log.Print("post error", err)
		c.JSON(400, nil)
		return
	}

	// パスワードをエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(post.Password), 14)

	postUser := domain.User{
		Name:     post.Name,
		Email:    post.Email,
		Password: password,
	}

	db := interactor.DB.Connect()

	user, err := interactor.User.Create(db, postUser)
	if err != nil {
		log.Print("ユーザー作成に失敗しました")
		c.JSON(400, nil)
		return
	}

	// JWTトークンを取得
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		log.Print("トークンの取得に失敗しました")
		return domain.User{}, NewResultStatus(400, nil)
	}

	// Cookieをセット
	cookie := new(http.Cookie)
	cookie.Value = token

	c.SetSameSite(http.SameSiteNoneMode)

	if os.Getenv("ENV") == "local" {
		log.Println("cookieをセットする")
		c.SetCookie("jwt", cookie.Value, 3600, "/", "localhost", true, true)
	}

	if os.Getenv("ENV") == "production" {
		log.Println("productionでcookieをセットする")
		c.SetCookie("jwt", cookie.Value, 3600, "/", "repgram-api.net", true, true)
	}

	if err != nil {
		log.Print("クッキーのセットに失敗しました")
		return domain.User{}, NewResultStatus(400, nil)
	}

	fmt.Println(c.Cookie("jwt"))

	return user, NewResultStatus(200, nil)
}
