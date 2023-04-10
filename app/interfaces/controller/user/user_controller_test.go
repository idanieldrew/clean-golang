package user

import (
	_ "bytes"
	"clean-golang/app/usecase/dto/user"
	interact "clean-golang/app/usecase/interactor/user"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	var mockInteract = &interact.InteractMock{}

	userController := UserController{
		Interact: mockInteract,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/users", nil)
	res := []user.PublicResponse{
		{
			Name:      "test1",
			Email:     "test1@gmail.com",
			Phone:     "111111",
			Password:  "password",
			UpdatedAt: time.Now(),
		},
		{
			Name:      "test2",
			Email:     "test2@gmail.com",
			Phone:     "222222",
			Password:  "password",
			UpdatedAt: time.Now(),
		},
	}
	
	status := 200
	mockInteract.On("Index").Return(res, status)
	userController.Index(w, r)

	if w.Code != status {
		t.Errorf("Unexpected response status code: got %v, want %v", w.Code, status)
	}
}

func TestProblemIndex(t *testing.T) {
	var mockInteract = &interact.InteractMock{}
	userController := UserController{Interact: mockInteract}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/users", nil)
	res := []user.PublicResponse{
		{},
	}
	status := 500
	mockInteract.On("Index").Return(res, status)
	userController.Index(w, r)

	if w.Code != status {
		t.Errorf("Unexpected response status code: got %v, want %v", w.Code, status)
	}
}
