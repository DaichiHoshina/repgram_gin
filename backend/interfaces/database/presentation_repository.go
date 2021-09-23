package database

import (
	"errors"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationRepository struct{}

func (repo *PresentationRepository) FindAll(db *gorm.DB, paginate domain.Paginate) (presentations domain.Presentations, err error) {
	presentations = domain.Presentations{}

	page, _ := strconv.Atoi(paginate.Page)
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(paginate.Per)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	log.Println(page, pageSize)

	db.Order("created_at DESC").
		Preload("User").
		Preload("Likes").
		Find(&presentations).
		Offset(offset).
		Limit(pageSize)

	return presentations, nil
}

func (repo *PresentationRepository) FindByID(db *gorm.DB, id int) (presentation domain.Presentation, err error) {
	presentation = domain.Presentation{}
	if id <= 0 {
		log.Println("IDがありません")
		return domain.Presentation{}, errors.New("IDがありません")
	}
	if result := db.First(&presentation, id); result.Error != nil {
		log.Println(presentation)
		return domain.Presentation{}, errors.New("投稿が取得出来ませんでした")
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
	presentation = domain.Presentation{}
	if id <= 0 {
		log.Println("IDがありません")
		return domain.Presentation{}, errors.New("IDがありません")
	}
	if result := db.Delete(&presentation, id); result.Error != nil {
		return domain.Presentation{}, errors.New("投稿が削除出来ませんでした")
	}
	return presentation, nil
}
