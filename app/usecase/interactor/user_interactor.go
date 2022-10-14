package interactor

import (
	"clean-golang/app/infrastructure/logger"
	"clean-golang/app/usecase/dto/user"
	"clean-golang/app/usecase/repository"
	"net/http"
)

type UserInteract struct {
	UserRepository repository.UserRepository
}

func (u *UserInteract) Index() ([]user.PublicResponse, int) {
	// Repository
	users, err := u.UserRepository.All()
	if err != nil {
		logger.Error("problem")
		return nil, http.StatusInternalServerError
	}

	r := user.UserResponse{}
	res := r.Public(users)
	return res, http.StatusOK
}
