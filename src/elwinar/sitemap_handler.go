package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sourcegraph/sitemap"
)

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
