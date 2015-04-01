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
			Name: "port,p",
			Usage: "set the port to listen on",
			EnvVar: "ELWINAR_PORT",
		},
		cli.IntFlag{
			Name: "timeout,t",
			Value: 1,
			Usage: "timeout for graceful shutdown",
			EnvVar: "ELWINAR_TIMEOUT",
		},
		cli.StringFlag{
			Name: "base,u",
			Value: "http://localhost:8080",
			Usage: "base url of the website",
			EnvVar: "ELWINAR_BASE_URL",
		},
	}
	
	app.Before = Bootstrap
	app.Action = Run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
