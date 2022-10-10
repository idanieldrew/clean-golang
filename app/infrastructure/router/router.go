package router

import (
	"clean-golang/app/infrastructure/router/user"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// register user route
	user.Router(r)
	return r
}
