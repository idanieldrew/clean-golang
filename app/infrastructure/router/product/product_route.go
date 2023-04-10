package product

import (
	"clean-golang/app/interfaces/controller/product"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	p := r.PathPrefix("/v1/products").Subrouter()

	pc := product.NewProduct()

	// store product
	p.HandleFunc("", pc.Store).Methods("POST")
	// find by slug
	p.HandleFunc("/{slug}", pc.FindBySlug).Methods("GET")
	// update
	p.HandleFunc("/{slug}", pc.Update).Methods("PATCH")
	// delete
	p.HandleFunc("/{slug}", pc.Destroy).Methods("DELETE")
}
