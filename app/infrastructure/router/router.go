package router

import (
	"clean-golang/app/infrastructure/router/product"
	"clean-golang/app/infrastructure/router/user"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	p := r.PathPrefix("/api").Subrouter()

	// register user route
	user.Router(p)
	// register product route
	product.Router(p)

	return r
}
