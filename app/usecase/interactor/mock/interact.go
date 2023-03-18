package mock

import (
	user_request "clean-golang/app/interfaces/request/user"
	"clean-golang/app/usecase/dto/user"
	"github.com/stretchr/testify/mock"
)

type InteractMock struct {
	mock.Mock
}

func (i *InteractMock) Index() ([]user.PublicResponse, int) {
	args := i.Called()
	return args.Get(0).([]user.PublicResponse), args.Get(1).(int)
}

func (i *InteractMock) Register(req *user_request.Request) (int, string) {
	//TODO implement me
	panic("implement me")
}
