package user

import (
	"clean-golang/app/entities"
	user_request "clean-golang/app/interfaces/request/user"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (u *MockRepo) All() (entities.Users, error) {
	args := u.Called()
	return args.Get(0).(entities.Users), args.Error(1)
}

func (u *MockRepo) Register(req *user_request.Request) error {
	return nil
}

func (u *MockRepo) CountMail(mail string) int {
	return 0
}

func (u *MockRepo) CountPhone(mail string) int {
	return 0
}
