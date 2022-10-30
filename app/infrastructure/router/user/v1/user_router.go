package user

import (
	"clean-golang/app/interfaces/controller"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	m := r.PathPrefix("/api").Subrouter()
	u := m.PathPrefix("/v1/users").Subrouter()

	uc := controller.New()

	// all users
	u.HandleFunc("", uc.Index).Methods("GET")
}
