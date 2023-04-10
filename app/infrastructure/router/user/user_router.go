package user

import (
	"clean-golang/app/interfaces/controller/user"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	u := r.PathPrefix("/v1/users").Subrouter()

	uc := user.New()

	// all users
	u.HandleFunc("", uc.Index).Methods("GET")

	u.HandleFunc("/register", uc.Register).Methods("POST")
}
