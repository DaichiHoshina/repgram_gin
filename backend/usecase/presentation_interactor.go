package usecase

import (
	"fmt"

	"github.com/DaichiHoshina/repgram_gin/backend/domain"
)

type PresentationInteractor struct {
	DB           DBRepository
	Presentation PresentationRepository
	S3           S3Repository
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

func (interactor *PresentationInteractor) PresentationCreate(c Context, awsS3 interface{}) (presentation domain.Presentation, resultStatus *ResultStatus) {
	var (
		err   error
		awsS3 AwsS3
		url   string
	)

	// 画像投稿処理
	upload_file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, err)
		return
	}

	src, err := upload_file.Open()
	if err != nil {
		c.JSON(400, err)
	}
	defer src.Close()

	awsS3 = domain.NewAwsS3()

	url, err = awsS3.UploadTest(src, upload_file.Filename, "png")

	if err != nil {
		fmt.Print(err.Error())
		c.JSON(400, err)
	}

	title := c.FormValue("title")
	userId := c.FormValue("user_id")
	discription := c.FormValue("discription")

	post := new(domain.Presentation)
	if err := c.Bind(post); err != nil {
		c.JSON(400, err)
	}
	postPresentation := domain.Presentation{
		Title:       title,
		UserID:      userId,
		Discription: discription,
		Image:       url,
	}

	user, res := controller.Interactor.PresentationCreate(postPresentation)
	if res.Error != nil {
		c.JSON(res.StatusCode, nil)
		return
	}
	c.JSON(res.StatusCode, user)
}
