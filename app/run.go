package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/stretchr/graceful"
	"log"
	"net/http"
	"time"
)

func Run(context *cli.Context) {
	err := run(context.Int("port"), context.String("secret"))
	if err != nil {
		log.Fatalln("unable to start:", err)
	}
}

func run(port int, secret string) error {
	r := httprouter.New()
	r.GET("/", IndexHandler)
	r.GET("/read", ReadHandler)
	r.GET("/article/:slug", ArticleHandler)
	r.GET("/article/:slug/edit", Auth(EditArticleHandler))
	r.POST("/article/:slug/edit", Auth(EditArticleFormHandler))
	r.GET("/article/:slug/delete", DeleteArticleHandler)
	r.GET("/article/:slug/publish", PublishArticleHandler)
	r.GET("/article/:slug/unpublish", UnpublishArticleHandler)
	r.GET("/login", LoginHandler)
	r.POST("/login", LoginFormHandler)
	r.GET("/logout", LogoutHandler)
	r.GET("/sitemap.xml", SitemapHandler)
	r.GET("/write", Auth(WriteHandler))
	r.POST("/write", Auth(WriteFormHandler))

	n := negroni.New()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewStatic(http.Dir("public")))
	n.Use(sessions.Sessions("elwinar", cookiestore.New([]byte(secret))))
	n.UseHandler(r)

	s := &graceful.Server{
		Timeout: 1 * time.Second,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: n,
		},
	}

	err := s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
