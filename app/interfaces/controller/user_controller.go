package controller

import (
	myMysql "clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/logger"
	"clean-golang/app/infrastructure/response"
	repo "clean-golang/app/interfaces/repository"
	"clean-golang/app/interfaces/request/user"
	"clean-golang/app/usecase/interactor"
	"encoding/json"
	"io"
	"net/http"
)

type UserController struct {
	Interact interactor.UserInteract
}

var (
	rp response.ResponseS
)

func Repo() *repo.UserRepository {
	return &repo.UserRepository{
		Connection: myMysql.Db,
		Cache:      redis.Rdb,
	}
}

func New() *UserController {
	return &UserController{
		Interact: interactor.UserInteract{
			UserRepository: Repo(),
		},
	}
}

func (u *UserController) Index(w http.ResponseWriter, r *http.Request) {
	res, status := u.Interact.Index()

	rp.Res(w, status, res)
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	req := new(user.Request)
	err := json.Unmarshal(body, req)
	if err != nil {
		logger.Error(err.Error())
		rp.Res(w, http.StatusInternalServerError, nil)
		return
	}

	validationErr := req.Validation()

	if validationErr != nil {
		logger.Error(validationErr.Error())
		rp.Res(w, http.StatusUnprocessableEntity, validationErr.Error())
		return
	}
	status, msg := u.Interact.Register(req)

	rp.Res(w, status, msg)
}
