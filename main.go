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
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/mattn/go-sqlite3"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/stretchr/graceful"
)

// Configuration holds the configuration for the website: which port to listen on,
// which database to use, etc.
type Configuration struct {
	Port     int
	Domain   string
	Database string
	Secret   string
	Password string
	Public   string
	GID      string
}

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

	// Add the front-office handlers.
	router.GET("/", IndexHandler)
	router.GET("/read", ArticleListHandler)
	router.GET("/article/:slug", ArticleViewHandler)
	router.GET("/fortune", QuoteFortuneHandler)
	router.GET("/sitemap.xml", SitemapHandler)

	// Add the back-office handlers.
	router.GET("/login", UserLoginHandler)
	router.POST("/login", UserAuthenticateHandler)
	router.GET("/logout", UserLogoutHandler)
	router.GET("/write", ArticleWriteHandler)
	router.POST("/write", ArticleCreateHandler)
	router.GET("/article/:slug/edit", ArticleEditHandler)
	router.POST("/article/:slug/edit", ArticleUpdateHandler)
	router.GET("/article/:slug/delete", ArticleDeleteHandler)
	router.GET("/article/:slug/publish", ArticlePublishHandler)
	router.GET("/article/:slug/unpublish", ArticleUnpublishHandler)

	// Initialize the server middleware stack.
	stack := negroni.New()
	stack.Use(gzip.Gzip(gzip.DefaultCompression))
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
