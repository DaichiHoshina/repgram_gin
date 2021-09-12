package database

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationRepository struct{}

func (repo *PresentationRepository) FindAll(db *gorm.DB) (presentation domain.Presentations, err error) {
	presentation = domain.Presentations{}
	db.Model(&presentation).
		Order("created_at DESC").
		// Preload("User").
		// Preload("Likes").
		Find(&presentation)

	fmt.Println(presentation)
	return presentation, nil
}

func (repo *PresentationRepository) FindByID(db *gorm.DB, id int) (presentation domain.Presentation, err error) {
	presentation = domain.Presentation{}
	db.First(&presentation, id)
	if presentation.ID <= 0 {
		return domain.Presentation{}, errors.New("presentation is not found")
	}
	fmt.Println(presentation)
	return presentation, nil
}
