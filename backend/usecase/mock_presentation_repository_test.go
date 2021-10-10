package usecase

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DaichiHoshina/repgram_gin/backend/domain"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func GetDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestFindByID(t *testing.T) {
	gdb, _, _ := GetDBMock()

	var expected domain.Presentation
	var err error

	// mockのコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成する
	mockApiClinet := NewMockPresentationRepository(ctrl)

	// 作成したmockに対して期待する呼び出しと返り値を定義する
	mockApiClinet.EXPECT().FindByID(gdb, 1).Return(expected, err)

	d := &presentationUsecase{}
	d.presentationRepo = mockApiClinet

	// mockを利用してtodoUsecase.View()をテストする
	res, err := d.presentationRepo.FindByID(gdb, 1)
	if err != nil {
		t.Error("Actual FindByID() is not same as expected")
	}

	assert.Equal(t, res, expected)
}

func TestAll(t *testing.T) {
	gdb, _, _ := GetDBMock()

	var expected domain.Presentations
	var err error

	paginate := domain.Paginate{
		Page: 1,
		Per:  6,
	}

	query := "test"

	// mockのコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成する
	mockApiClinet := NewMockPresentationRepository(ctrl)

	// 作成したmockに対して期待する呼び出しと返り値を定義する
	mockApiClinet.EXPECT().FindAll(gdb, paginate, query).Return(expected, err)

	d := &presentationUsecase{}
	d.presentationRepo = mockApiClinet

	// mockを利用してtodoUsecase.FindAll()をテストする
	res, err := d.presentationRepo.FindAll(gdb, paginate, query)
	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	assert.Equal(t, res, expected)
}

func TestCreate(t *testing.T) {
	gdb, _, _ := GetDBMock()

	var expected domain.Presentation
	var err error

	postPresentation := domain.Presentation{
		Title:       "タイトル",
		UserID:      "1",
		Discription: "文章です",
		Image:       "url_test.jp",
	}

	// mockのコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成する
	mockApiClinet := NewMockPresentationRepository(ctrl)

	// 作成したmockに対して期待する呼び出しと返り値を定義する
	mockApiClinet.EXPECT().Create(gdb, postPresentation).Return(expected, err)

	d := &presentationUsecase{}
	d.presentationRepo = mockApiClinet

	// mockを利用してtodoUsecase.Create()をテストする
	res, err := d.presentationRepo.Create(gdb, postPresentation)
	if err != nil {
		t.Error("Actual Create() is not same as expected")
	}

	assert.Equal(t, res, expected)
}

func TestUpdate(t *testing.T) {
	gdb, _, _ := GetDBMock()

	var expected domain.Presentation
	var err error

	modelPresentation := domain.Presentation{}

	postPresentation := domain.Presentation{
		Title:       "タイトル",
		UserID:      "1",
		Discription: "文章です",
		Image:       "url_test.jp",
	}

	// mockのコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成する
	mockApiClinet := NewMockPresentationRepository(ctrl)

	// 作成したmockに対して期待する呼び出しと返り値を定義する
	mockApiClinet.EXPECT().Update(gdb, postPresentation, modelPresentation).Return(expected, err)

	d := &presentationUsecase{}
	d.presentationRepo = mockApiClinet

	// mockを利用してtodoUsecase.Update()をテストする
	res, err := d.presentationRepo.Update(gdb, postPresentation, modelPresentation)
	if err != nil {
		t.Error("Actual Update() is not same as expected")
	}

	assert.Equal(t, res, expected)
}

func TestDelete(t *testing.T) {
	gdb, _, _ := GetDBMock()

	var expected domain.Presentation
	var err error

	// mockのコントローラを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成する
	mockApiClinet := NewMockPresentationRepository(ctrl)

	// 作成したmockに対して期待する呼び出しと返り値を定義する
	mockApiClinet.EXPECT().Delete(gdb, 1).Return(expected, err)

	d := &presentationUsecase{}
	d.presentationRepo = mockApiClinet

	// mockを利用してtodoUsecase.Delete()をテストする
	res, err := d.presentationRepo.Delete(gdb, 1)
	if err != nil {
		t.Error("Actual Delete() is not same as expected")
	}

	assert.Equal(t, res, expected)
}
