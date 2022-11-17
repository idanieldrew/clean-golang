package repository

import (
	"clean-golang/app/entities"
	user_request "clean-golang/app/interfaces/request/user"
)

type UserRepository interface {
	All() (entities.Users, error)
	Register(req *user_request.Request) error
}
