package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/graceful"
	"log"
	"net/http"
	"time"
)

func Run(context *cli.Context) {
	err := run(context.Int("port"), context.Int("timeout"))
	if err != nil {
		log.Fatalln("unable to start:", err)
	}
}

func run(port, timeout int) error {
	r := httprouter.New()
	r.GET("/", IndexHandler)
	r.GET("/read", ReadHandler)
	r.GET("/a/:slug", ArticleHandler)
	
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir("public/")))
	n.UseHandler(r)
	
	s := &graceful.Server{
		Timeout: time.Duration(timeout) * time.Second,
		Server: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
			Handler: n,
		},
	}
	
	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	
	return nil
}
