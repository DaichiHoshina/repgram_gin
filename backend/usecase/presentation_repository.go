package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Presentation, err error)
}
