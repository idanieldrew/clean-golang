package user

import (
	user_request "clean-golang/app/interfaces/request/user"
	"clean-golang/app/usecase/dto/user"
)

type Interact interface {
	Index() ([]user.PublicResponse, int)
	Register(req *user_request.Request) (int, string)
}
