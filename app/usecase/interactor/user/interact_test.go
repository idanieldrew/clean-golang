package user

import (
	"clean-golang/app/entities"
	"clean-golang/app/usecase/repository/user"
	"errors"
	"testing"
	"time"
)

func TestInteractIndex(t *testing.T) {
	repo := &user.MockRepo{}
	userInteract := UserInteract{UserRepository: repo}
	res := entities.Users{
		{
			Id:              1,
			Name:            "test",
			Email:           "test@test.test",
			Phone:           "11111",
			EmailVerifiedAt: "",
			Password:        "password",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	repo.On("All").Return(res, nil)

	status := 200
	_, expect := userInteract.Index()
	if expect != status {
		t.Errorf("%d is  not equal with %d", expect, status)
	}
}

func TestExistProblemInRepoLayer(t *testing.T) {
	repo := &user.MockRepo{}
	userInteract := UserInteract{UserRepository: repo}
	res := entities.Users{}
	repo.On("All").Return(res, errors.New("problem in scan"))

	status := 500
	_, expect := userInteract.Index()
	if expect != status {
		t.Errorf("%d is  not equal with %d", expect, status)
	}
}
