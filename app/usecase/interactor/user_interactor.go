package interactor

import (
	"clean-golang/app/infrastructure/logger"
	user_request "clean-golang/app/interfaces/request/user"
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

func (u *UserInteract) Register(req *user_request.Request) (int, string) {
	// check validation usecase
	request := &user.Request{
		Email: req.Email,
		Phone: req.Phone,
	}
	if req := request.Validation(u.UserRepository); !req {
		return http.StatusUnprocessableEntity, "incorrect mail or phone"
	}

	// register user
	err := u.UserRepository.Register(req)
	if err != nil {
		return http.StatusInternalServerError, "server problem"
	}

	return http.StatusCreated, "success"
}
