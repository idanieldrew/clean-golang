package controller_test

import (
	_ "bytes"
	"clean-golang/app/interfaces/controller"
	"clean-golang/app/usecase/dto/user"
	"clean-golang/app/usecase/interactor/mock"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorrectIndex(t *testing.T) {
	var mockInteract = &mock.InteractMock{}
	userController := controller.UserController{Interact: mockInteract}
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
	var mockInteract = &mock.InteractMock{}
	userController := controller.UserController{Interact: mockInteract}
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
