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
		return domain.Presentation{}, errors.New("投稿が見つかりませんでした")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Create(db *gorm.DB, postPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	if result := db.Create(&postPresentation); result.Error != nil {
		return domain.Presentation{}, errors.New("投稿が作成出来ませんでした")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Update(db *gorm.DB, postPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	if result := db.Model(&presentation).Update(&postPresentation); result.Error != nil {
		return domain.Presentation{}, errors.New("投稿が更新出来ませんでした")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Delete(db *gorm.DB, id int) (presentation domain.Presentation, err error) {
	if result := db.First(&presentation, id).Delete(&presentation); result.Error != nil {
		return domain.Presentation{}, errors.New("投稿が削除出来ませんでした")
	}
	return presentation, nil
}
