package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationRepository interface {
	FindByID(db *gorm.DB, id int) (presentation domain.Presentation, err error)
	FindAll(db *gorm.DB, post domain.Paginate) (presentation domain.Presentations, err error)
	Create(db *gorm.DB, postPresentation domain.Presentation) (presentation domain.Presentation, err error)
	Update(db *gorm.DB, postPresentation domain.Presentation) (presentation domain.Presentation, err error)
	Delete(db *gorm.DB, id int) (presentation domain.Presentation, err error)
}
