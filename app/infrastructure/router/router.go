package router

import (
	"clean-golang/app/interfaces/controller"
	"fmt"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index)
	fmt.Print("ok")
}
