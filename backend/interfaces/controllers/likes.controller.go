package controllers

import (
	"fmt"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
	"github.com/DaichiHoshina/repgram_gin/backend/interfaces/database"
	"github.com/DaichiHoshina/repgram_gin/backend/usecase"
)

type LikesController struct {
	Interactor usecase.LikeInteractor
}

func NewLikesController(db database.DB) *LikesController {
	return &LikesController{
		Interactor: usecase.LikeInteractor{
			DB:   &database.DBRepository{DB: db},
			Like: &database.LikeRepository{},
		},
	}
}

func (controller *LikesController) Create(c Context) {
	post := new(domain.Like)
	if err := c.Bind(post); err != nil {
		c.JSON(400, nil)
	}

	fmt.Println(post.UserID, post.PresentationID)
	postLike := domain.Like{
		UserID:         post.UserID,
		PresentationID: post.PresentationID,
	}

	user, res := controller.Interactor.LikeCreate(postLike)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, user)
}

func (controller *LikesController) Delete(c Context) {

	post := new(domain.Like)
	if err := c.Bind(post); err != nil {
		c.JSON(400, nil)
	}

	fmt.Println(post.UserID, post.PresentationID)
	postLike := domain.Like{
		UserID:         post.UserID,
		PresentationID: post.PresentationID,
	}

	user, res := controller.Interactor.LikeDelete(postLike)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, user)
}
