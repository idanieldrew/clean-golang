package cmd

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Run() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"serve"},
				Usage:   "serve project",
				Action:  Serve,
			},
			{
				Name:    "migrate",
				Aliases: []string{"migrate"},
				Usage:   "migrate db",
				Action:  Migrate,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
