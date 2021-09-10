package database

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationRepository struct{}

func (repo *PresentationRepository) FindByID(db *gorm.DB, id int) (presentation domain.Presentation, err error) {
	presentation = domain.Presentation{}
	db.First(&presentation, id)
	if presentation.ID <= 0 {
		return domain.Presentation{}, errors.New("presentation is not found")
	}
	return presentation, nil
}
