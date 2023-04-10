package user

import (
	myMysql "clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/logger"
	"clean-golang/app/interfaces/controller"
	repo "clean-golang/app/interfaces/repository/user"
	"clean-golang/app/interfaces/request/user"
	Interact "clean-golang/app/usecase/interactor/user"
	"encoding/json"
	"io"
	"net/http"
)

type UserController struct {
	Interact Interact.Interact
}

func Repo() *repo.UserRepository {
	// mongo repo
	return &repo.UserRepository{
		Connection: myMysql.Db,
		Cache:      redis.Rdb,
	}
}

func New() *UserController {
	return &UserController{
		Interact: &Interact.UserInteract{
			UserRepository: Repo(),
		},
	}
}

func (u *UserController) Index(w http.ResponseWriter, r *http.Request) {
	res, status := u.Interact.Index()
	controller.Res.Res(w, status, res)
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	req := new(user.Request)
	err := json.Unmarshal(body, req)
	if err != nil {
		logger.Error(err.Error())
		controller.Res.Res(w, http.StatusInternalServerError, nil)
		return
	}

	validationErr := req.Validation()

	if validationErr != nil {
		logger.Error(validationErr.Error())
		controller.Res.Res(w, http.StatusUnprocessableEntity, validationErr.Error())
		return
	}
	status, msg := u.Interact.Register(req)

	controller.Res.Res(w, status, msg)
}
