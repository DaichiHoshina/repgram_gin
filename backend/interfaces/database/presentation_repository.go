package database

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationRepository struct{}

func (repo *PresentationRepository) FindAll(db *gorm.DB) (presentations domain.Presentations, err error) {
	presentations = domain.Presentations{}
	// db.Model(&presentation.Presentation).

	db.Order("created_at DESC").
		Preload("User").
		Preload("Likes").
		Find(&presentations)

	return presentations, nil
}

func (repo *PresentationRepository) FindByID(db *gorm.DB, id int) (presentation domain.Presentation, err error) {
	presentation = domain.Presentation{}
	db.First(&presentation, id)
	if presentation.ID <= 0 {
		return domain.Presentation{}, errors.New("presentation is not found")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Create(db *gorm.DB, postPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	if result := db.Create(&postPresentation); result.Error != nil {
		return domain.Presentation{}, errors.New("presentation can not create")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Update(db *gorm.DB, postPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	if result := db.Model(&presentation).Update(&postPresentation); result.Error != nil {
		return domain.Presentation{}, errors.New("presentation can not create")
	}
	return presentation, nil
}
