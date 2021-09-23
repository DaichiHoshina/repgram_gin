package database

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type LikeRepository struct{}

func (repo *LikeRepository) Create(db *gorm.DB, postLike domain.Like) (like domain.Like, err error) {
	if result := db.Create(&postLike); result.Error != nil {
		return domain.Like{}, errors.New("いいねが作成出来ませんでした")
	}
	return like, nil
}

func (repo *LikeRepository) Delete(db *gorm.DB, postLike domain.Like) (like domain.Like, err error) {

	like = domain.Like{
		UserID:         postLike.UserID,
		PresentationID: postLike.PresentationID,
	}

	db.First(&like, like)

	if result := db.Delete(&like); result.Error != nil {
		return domain.Like{}, errors.New("いいねが削除出来ませんでした")
	}
	return like, nil
}
