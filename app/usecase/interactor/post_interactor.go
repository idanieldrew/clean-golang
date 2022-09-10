package interactor

import "clean-golang/app/usecase/repository"

type UserInteract struct {
	UserRepository repository.UserRepository
}
