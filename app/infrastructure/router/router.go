package router

import (
	"clean-golang/app/interfaces/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index).Methods("GET")
	return r
}
