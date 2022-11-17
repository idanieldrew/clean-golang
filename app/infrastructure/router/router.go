package router

import (
	"clean-golang/app/infrastructure/router/user/v1"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	p := r.PathPrefix("/api").Subrouter()

	// register user route
	user.Router(p)
	return r
}
