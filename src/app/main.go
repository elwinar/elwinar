package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"github.com/sourcegraph/sitemap"
	"github.com/stretchr/graceful"
	"github.com/urfave/negroni"
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
	stack.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
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

// Index displays the home page of the website.
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	timezone, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		timezone = time.FixedZone("Europe/Paris", 1)
	}

	birthDate := time.Date(1990, time.November, 5, 0, 0, 0, 0, timezone)
	now := time.Now()

	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	render(w, r, "index", map[string]interface{}{
		"Title": "Passionate developer",
		"Age":   age,
	})
}

// Sitemap will display an XML map of the website.
func SitemapHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var articles []*Article

	// Get each article.
	err := database.Select(&articles, "SELECT id, title, slug, tagline, text, tags, is_published, created_at, updated_at, published_at FROM articles WHERE is_published = ? ORDER BY updated_at DESC", true)
	if err != nil {
		panic(err)
	}

	// Add the static pages.
	var urlset sitemap.URLSet
	urlset.URLs = []sitemap.URL{
		{
			Loc:        fmt.Sprintf("%s/", r.Host),
			ChangeFreq: sitemap.Yearly,
		},
		{
			Loc:        fmt.Sprintf("%s/read", r.Host),
			ChangeFreq: sitemap.Weekly,
		},
	}

	// Add the article pages.
	for _, a := range articles {
		urlset.URLs = append(urlset.URLs, sitemap.URL{
			Loc:        fmt.Sprintf("%s/article/%s", r.Host, a.Slug),
			LastMod:    &a.UpdatedAt,
			ChangeFreq: sitemap.Monthly,
		})
	}

	// Marshal the sitemap to XML.
	raw, err := sitemap.Marshal(&urlset)
	if err != nil {
		panic(err)
	}

	w.Write(raw)
}
