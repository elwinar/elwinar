package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/graceful"
)

var (
	configuration Configuration
	database      *sqlx.DB
)

func main() {
	// Load the configuration.
	err := envconfig.Process("app", &configuration)
	if err != nil {
		log.Fatalln("unable to read the configuration from env:", err)
	}

	// Open the database connection.
	database, err = sqlx.Connect("sqlite3", configuration.Database)
	if err != nil {
		log.Fatalln("unable to open the database:", err)
	}

	// Initialize the router.
	router := httprouter.New()

	if configuration.Debug {
		router.GET("/env", EnvironmentHandler)
		router.GET("/cfg", ConfigurationHandler)
	}

	// Add the front-office handlers.
	router.GET("/", IndexHandler)
	router.GET("/read", ArticleListHandler)
	router.GET("/articles/:slug", ArticleViewHandler)
	router.GET("/fortune", QuoteFortuneHandler)
	router.GET("/quotes/:id", QuoteViewHandler)
	router.GET("/sitemap.xml", SitemapHandler)

	// Add the back-office handlers.
	router.GET("/login", UserLoginHandler)
	router.POST("/login", UserAuthenticateHandler)
	router.GET("/logout", UserLogoutHandler)
	router.GET("/write", ArticleWriteHandler)
	router.POST("/write", ArticleCreateHandler)
	router.GET("/articles/:slug/edit", ArticleEditHandler)
	router.POST("/articles/:slug/edit", ArticleUpdateHandler)
	router.GET("/articles/:slug/delete", ArticleDeleteHandler)
	router.GET("/articles/:slug/publish", ArticlePublishHandler)
	router.GET("/articles/:slug/unpublish", ArticleUnpublishHandler)
	router.GET("/quotes", QuoteListHandler)

	// Initialize the server middleware stack.
	stack := negroni.New()
	stack.Use(negroni.NewRecovery())
	stack.Use(negroni.NewStatic(http.Dir(configuration.Public)))
	stack.Use(sessions.Sessions("elwinar", cookiestore.New([]byte(configuration.Secret))))
	stack.UseHandler(router)

	// Initialize the HTTP server.
	server := &graceful.Server{
		Timeout: 1 * time.Second,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", configuration.Port),
			Handler: stack,
		},
	}

	// Run the server.
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("unable to run the server:", err)
	}
}
