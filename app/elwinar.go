package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	var app = cli.NewApp()

	app.Name = "elwinar"
	app.Version = "4go-dev"
	app.Author = "Romain Baugue"
	app.Email = "romain.baugue@elwinar.com"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Usage:  "listening port",
			EnvVar: "ELWINAR_PORT",
		},
		cli.StringFlag{
			Name:   "base",
			Usage:  "base url",
			EnvVar: "ELWINAR_BASE",
		},
		cli.StringFlag{
			Name:   "database",
			Usage:  "sqlite database file",
			EnvVar: "ELWINAR_DATABASE",
		},
		cli.StringFlag{
			Name:   "secret",
			Usage:  "encryption secret",
			EnvVar: "ELWINAR_SECRET",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "administrative password",
			EnvVar: "ELWINAR_PASSWORD",
		},
		cli.StringFlag{
			Name:   "views",
			Usage:  "views repository",
			EnvVar: "ELWINAR_VIEWS",
		},
	}

	app.Before = Bootstrap
	app.Action = Run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
