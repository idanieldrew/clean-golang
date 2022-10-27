package controller

import (
	myMysql "clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/database/redis"
	"clean-golang/app/infrastructure/response"
	repo "clean-golang/app/interfaces/repository"
	"clean-golang/app/usecase/interactor"
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

	rp := response.ResponseS{}
	rp.Res(w, status, res)
}
