package controller

import (
	"clean-golang/app/usecase/interactor"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	interactor.Index()
}
