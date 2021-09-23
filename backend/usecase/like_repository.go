package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type LikeRepository interface {
	Create(db *gorm.DB, postLike domain.Like) (presentation domain.Like, err error)
	Delete(db *gorm.DB, postLike domain.Like) (presentation domain.Like, err error)
}
