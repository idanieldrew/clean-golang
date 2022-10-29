package user

import (
	"clean-golang/app/interfaces/controller"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	u := r.PathPrefix("/users").Subrouter()

	uc := controller.New()

	// all users
	u.HandleFunc("", uc.Index).Methods("GET")
}
