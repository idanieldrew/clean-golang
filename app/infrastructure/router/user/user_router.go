package user

import (
	"clean-golang/app/interfaces/controller"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	u := r.PathPrefix("/users").Subrouter()
	// all users
	u.HandleFunc("", controller.Index).Methods("GET")
}
