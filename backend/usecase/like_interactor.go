package usecase

import (
	"log"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type LikeInteractor struct {
	DB   DBRepository
	Like LikeRepository
}

func (interactor *LikeInteractor) LikeCreate(postLike domain.Like) (like domain.Like, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	like, err := interactor.Like.Create(db, postLike)
	if err != nil {
		log.Println("いいねが作成出来ませんでした")
		return domain.Like{}, NewResultStatus(400, err)
	}
	return like, NewResultStatus(200, nil)
}

func (interactor *LikeInteractor) LikeDelete(postLike domain.Like) (like domain.Like, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	like, err := interactor.Like.Delete(db, postLike)
	if err != nil {
		log.Println("いいねが削除出来ませんでした")
		return domain.Like{}, NewResultStatus(400, err)
	}
	return like, NewResultStatus(200, nil)
}
