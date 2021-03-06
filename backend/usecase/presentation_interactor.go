package usecase

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationInteractor struct {
	DB           DBRepository
	Presentation PresentationRepository
}

func (interactor *PresentationInteractor) Presentations(c Context) (presentation domain.Presentations, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	page := c.Query("page")
	per := c.Query("per")
	searchQuery := c.Query("query")

	page_int, _ := strconv.Atoi(page)
	per_int, _ := strconv.Atoi(per)

	paginate := domain.Paginate{
		Page: page_int,
		Per:  per_int,
	}

	presentations, err := interactor.Presentation.FindAll(db, paginate, searchQuery)
	if err != nil {
		log.Println("投稿が取得出来ませんでした")
		return domain.Presentations{}, NewResultStatus(400, err)
	}
	return presentations, NewResultStatus(200, nil)
}

func (interactor *PresentationInteractor) PresentationByID(id int) (presentation domain.PresentationForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundPresentation, err := interactor.Presentation.FindByID(db, id)
	if err != nil {
		log.Println("投稿が見つかりませんでした")
		return domain.PresentationForGet{}, NewResultStatus(400, err)
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
		log.Println("画像アップロードに失敗しました。")
		return domain.Presentation{}, NewResultStatus(400, err)
	}

	src, err := upload_file.Open()
	if err != nil {
		log.Println("画像アップロードに失敗しました。")
		return domain.Presentation{}, NewResultStatus(400, err)
	}
	defer src.Close()

	awsS3 = domain.NewAwsS3()

	url, err = awsS3.UploadTest(src, upload_file.Filename, "png")
	if err != nil {
		fmt.Print(err.Error())
		log.Println("S３のアップロードに失敗しました。")
		return domain.Presentation{}, NewResultStatus(400, err)
	}

	title, _ := c.GetPostForm("title")
	userId, _ := c.GetPostForm("user_id")
	discription, _ := c.GetPostForm("discription")

	post := new(domain.Presentation)
	if err := c.Bind(post); err != nil {
		log.Println("paramの取得に失敗しました")
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

		presentation, err := interactor.Presentation.Update(db, postPresentation, presentation)
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

func (interactor *PresentationInteractor) PresentationDelete(id int) (presentation domain.Presentation, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	presentation, err := interactor.Presentation.Delete(db, id)
	if err != nil {
		log.Println("投稿の削除に失敗しました。")
		return domain.Presentation{}, NewResultStatus(400, err)
	}
	return presentation, NewResultStatus(200, nil)
}
