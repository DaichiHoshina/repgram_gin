package database

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationRepository struct{}

func (repo *PresentationRepository) FindAll(db *gorm.DB, paginate domain.Paginate, query string) (presentations domain.Presentations, err error) {
	presentations = domain.Presentations{}

	page := paginate.Page
	pageSize := paginate.Per
	// 先頭いくつスキップするかを取得
	offset := (page - 1) * pageSize

	log.Println(query + "で検索")

	db.Debug().Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Preload("User").
		Preload("Likes").
		Where("discription LIKE ?", "%"+query+"%").
		Find(&presentations)

	return presentations, nil
}

func (repo *PresentationRepository) FindByID(db *gorm.DB, id int) (presentation domain.Presentation, err error) {
	presentation = domain.Presentation{}
	if id <= 0 {
		log.Println("IDがありません")
		return domain.Presentation{}, errors.New("IDがありません")
	}
	if result := db.First(&presentation, id); result.Error != nil {
		return domain.Presentation{}, errors.New("投稿が取得出来ませんでした")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Create(db *gorm.DB, postPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	if result := db.Create(&postPresentation); result.Error != nil {
		log.Println("投稿が作成出来ませんでした")
		return domain.Presentation{}, errors.New("投稿が作成出来ませんでした")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Update(db *gorm.DB, postPresentation domain.Presentation, modelPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	if result := db.Model(&modelPresentation).Update(&postPresentation); result.Error != nil {
		log.Println("投稿が更新出来ませんでした")
		return domain.Presentation{}, errors.New("投稿が更新出来ませんでした")
	}
	return presentation, nil
}

func (repo *PresentationRepository) Delete(db *gorm.DB, id int) (presentation domain.Presentation, err error) {
	presentation = domain.Presentation{
		ID: id,
	}

	if id <= 0 {
		log.Println("IDがありません")
		return domain.Presentation{}, errors.New("IDがありません")
	}
	db.First(&presentation, id)

	if result := db.Delete(&presentation); result.Error != nil {
		log.Println("投稿が削除出来ませんでした")
		return domain.Presentation{}, errors.New("投稿が削除出来ませんでした")
	}
	return presentation, nil
}
