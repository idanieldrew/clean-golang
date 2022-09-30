package controller

import (
	mysql2 "clean-golang/app/infrastructure/database/mysql"
	repo "clean-golang/app/interfaces/repository"
	"clean-golang/app/usecase/interactor"
	"fmt"
	"net/http"
)

type UserController struct {
	Interact interactor.UserInteract
}

func New() *UserController {
	return &UserController{
		Interact: interactor.UserInteract{
			UserRepository: &repo.UserRepository{
				Connection: mysql2.Db,
			},
		},
	}
}

func (u *UserController) Index(w http.ResponseWriter, r *http.Request) {
	interact := u.Interact.Index()
	fmt.Println(interact)
}
