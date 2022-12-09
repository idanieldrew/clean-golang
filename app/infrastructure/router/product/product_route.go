package product

import (
	"clean-golang/app/interfaces/controller"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	p := r.PathPrefix("/v1/products").Subrouter()

	pc := controller.NewProduct()

	// store product
	p.HandleFunc("", pc.Store).Methods("POST")
}
