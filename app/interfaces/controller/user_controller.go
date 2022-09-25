package controller

import (
	"clean-golang/app/usecase/interactor"
	"fmt"
	"net/http"
)

type UserController struct {
	Interact interactor.UserInteract
}

func New() *UserController {
	return &UserController{Interact: interactor.UserInteract{}}
}

func (u *UserController) Index(w http.ResponseWriter, r *http.Request) {
	interact := u.Interact.Index()
	fmt.Println(interact)
}
