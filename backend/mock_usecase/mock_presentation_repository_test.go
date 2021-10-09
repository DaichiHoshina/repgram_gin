package mock

import (
	"testing"

	mock_usecase "github.com/DaichiHoshina/repgram_gin/backend/mock_usecase"
	"github.com/golang/mock/gomock"
)

func TestSample(t *testing.T) {
	// mockのコントローラを作成します
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成します
	mockApiClinet := mock_usecase.NewMockPresentationRepository(ctrl)
	// 作成したmockに対して期待する呼び出しと返り値を定義します
	// EXPECT()では呼び出されたかどうか
	// Request()ではそのメソッド名が指定した引数で呼び出されたかどうか
	// Return()では返り値を指定します
	mockApiClinet.EXPECT().Request("bar").Return("bar", nil)

	d := &DataRegister{}
	d.client = mockCpiClinet // mockを登録
	expected := "bar"

	res, err := d.Register(expected)
	if err != nil {
		t.Fatal("Register error!", err)
	}
	if res != expected {
		t.Fatal("Value does not match.")
	}
}

// func TestView(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	var expected []*domain.Presentation
// 	var err error

// 	mockSample := NewMockPresentationRepository(ctrl)
// 	mockSample.EXPECT().FindByID().Return(expected, err)

// 	// mockを利用してtodoUsecase.View()をテストする
// 	todoUsecase := controllers.NewPresentationsController().NewPresentationUsecase(mockSample)
// 	result, err := todoUsecase.View()

// 	if err != nil {
// 		t.Error("Actual FindAll() is not same as expected")
// 	}

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Actual FindAll() is not same as expected")
// 	}

// }

// func TestSearch(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	var expected []*domain.Presentation
// 	var err error
// 	word := "test"

// 	mockSample := NewMockPresentationRepository(ctrl)
// 	mockSample.EXPECT().Find(word).Return(expected, err)

// 	// mockを利用してtodoUsecase.Search(word string)をテストする
// 	todoUsecase := usecase.NewPresentationUsecase(mockSample)
// 	result, err := todoUsecase.Search(word)

// 	if err != nil {
// 		t.Error("Actual Find(word string) is not same as expected")
// 	}

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Actual Find(word string) is not same as expected")
// 	}

// }

// func TestAdd(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	var expected *domain.Presentation
// 	var err error

// 	mockSample := NewMockPresentationRepository(ctrl)
// 	mockSample.EXPECT().Create(expected).Return(expected, err)

// 	// mockを利用してtodoUsecase.Add(todo *domain.Presentation)をテストする
// 	todoUsecase := usecase.NewPresentationUsecase(mockSample)
// 	err = todoUsecase.Add(expected)

// 	if err != nil {
// 		t.Error("Actual Find(word string) is not same as expected")
// 	}

// }

// func TestEdit(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	var expected *domain.Presentation
// 	var err error

// 	mockSample := NewMockPresentationRepository(ctrl)
// 	mockSample.EXPECT().Update(expected).Return(expected, err)

// 	// mockを利用してtodoUsecase.Edit(todo *domain.Presentation)をテストする
// 	todoUsecase := usecase.NewPresentationUsecase(mockSample)
// 	err = todoUsecase.Edit(expected)

// 	if err != nil {
// 		t.Error("Actual Find(word string) is not same as expected")
// 	}

// }
