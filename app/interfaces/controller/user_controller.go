package controller

import (
	"clean-golang/app/usecase/interactor"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	u := interactor.Index()
	fmt.Println(u)
}
