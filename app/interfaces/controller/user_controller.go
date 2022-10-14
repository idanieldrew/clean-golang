package controller

import (
	myMysql "clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/logger"
	repo "clean-golang/app/interfaces/repository"
	"clean-golang/app/usecase/interactor"
	"encoding/json"
	"net/http"
)

type UserController struct {
	Interact interactor.UserInteract
}

func New() *UserController {
	return &UserController{
		Interact: interactor.UserInteract{
			UserRepository: &repo.UserRepository{
				Connection: myMysql.Db,
				Cache:      redis.Rdb,
			},
		},
	}
}

func (u *UserController) Index(w http.ResponseWriter, r *http.Request) {
	res, status := u.Interact.Index()

	u.response(w, status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Error("problem in response user")
		return
	}
}

func (u *UserController) response(w http.ResponseWriter, s int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	return
}
