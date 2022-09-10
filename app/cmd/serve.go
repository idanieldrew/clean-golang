package cmd

import (
	"clean-golang/app/infrastructure/router"
	"fmt"
	"github.com/urfave/cli/v2"
)

func Serve(cCtx *cli.Context) error {
	// register routes
	router.Router()
	fmt.Println("salam")
	return nil
}
