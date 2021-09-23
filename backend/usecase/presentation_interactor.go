package usecase

import (
	"fmt"
	"log"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationInteractor struct {
	DB           DBRepository
	Presentation PresentationRepository
}

func (interactor *PresentationInteractor) Presentations() (presentation domain.Presentations, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	presentations, err := interactor.Presentation.FindAll(db)
	if err != nil {
		return domain.Presentations{}, NewResultStatus(404, err)
	}
	return presentations, NewResultStatus(200, nil)
}

func (interactor *PresentationInteractor) PresentationByID(id int) (presentation domain.PresentationForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundPresentation, err := interactor.Presentation.FindByID(db, id)
	if err != nil {
		return domain.PresentationForGet{}, NewResultStatus(404, err)
	}
	presentation = foundPresentation.BuildForGet()
	return presentation, NewResultStatus(200, nil)
}

func (interactor *PresentationInteractor) PresentationCreate(c Context) (presentation domain.Presentation, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	var (
		err   error
		awsS3 *domain.AwsS3
		url   string
	)

	// 画像投稿処理
	upload_file, err := c.FormFile("file")
	if err != nil {
		return domain.Presentation{}, NewResultStatus(400, err)
	}

	src, err := upload_file.Open()
	if err != nil {
		return domain.Presentation{}, NewResultStatus(400, err)
	}
	defer src.Close()

	awsS3 = domain.NewAwsS3()

	url, err = awsS3.UploadTest(src, upload_file.Filename, "png")
	if err != nil {
		fmt.Print(err.Error())
		return domain.Presentation{}, NewResultStatus(400, err)
	}

	title, _ := c.GetPostForm("title")
	userId, _ := c.GetPostForm("user_id")
	discription, _ := c.GetPostForm("discription")

	post := new(domain.Presentation)
	if err := c.Bind(post); err != nil {
		return domain.Presentation{}, NewResultStatus(400, err)
	}

	postPresentation := domain.Presentation{
		Title:       title,
		UserID:      userId,
		Discription: discription,
		Image:       url,
	}

	presentation, err = interactor.Presentation.Create(db, postPresentation)
	if err != nil {
		log.Println("投稿作成に失敗しました")
		c.JSON(400, "投稿作成に失敗しました")
		return
	}
	return presentation, NewResultStatus(200, nil)
}

func (interactor *PresentationInteractor) PresentationUpdate(c Context) (presentation domain.Presentation, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	if id := c.Param("id"); id != "" {

		db.First(&presentation, id)
		post := new(domain.Presentation)

		if err := c.Bind(post); err != nil {
			log.Println("投稿更新に失敗しました")
			c.JSON(400, "投稿更新に失敗しました")
		}

		postPresentation := domain.Presentation{
			Discription: post.Discription,
		}

		presentation, err := interactor.Presentation.Update(db, postPresentation)
		if err != nil {
			log.Println("投稿更新に失敗しました")
			c.JSON(400, "投稿更新に失敗しました")
		}
		return presentation, NewResultStatus(200, nil)
	}
	log.Println("IDが取得できませんでした")
	c.JSON(400, "IDが取得できませんでした")
	return
}
