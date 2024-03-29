package cmd

import (
	"clean-golang/app/infrastructure/router"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
)

func Serve(cCtx *cli.Context) error {
	// register routes
	r := router.Router()
	log.SetPrefix("[SUCCESS] ")
	log.Println("Serving project http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
	return nil
}
