package interactor

import (
	"clean-golang/app/entities"
	"clean-golang/app/infrastructure/logger"
	"clean-golang/app/usecase/repository"
)

type UserInteract struct {
	UserRepository repository.UserRepository
}

func (u *UserInteract) Index() entities.Users {
	users, err := u.UserRepository.All()
	if err != nil {
		logger.Error("problem")
		return nil
	}
	return users
}
