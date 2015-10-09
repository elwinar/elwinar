package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sourcegraph/sitemap"
)

func SitemapHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var articles []*Article

	err := database.Select(&articles, "SELECT id, title, slug, tagline, text, tags, is_published, created_at, updated_at, published_at FROM articles WHERE is_published = ? ORDER BY updated_at DESC", true)
	if err != nil {
		panic(err)
	}

	var urlset sitemap.URLSet
	urlset.URLs = []sitemap.URL{
		{
			Loc:        fmt.Sprintf("%s/", configuration.Base),
			ChangeFreq: sitemap.Yearly,
		},
		{
			Loc:        fmt.Sprintf("%s/read", configuration.Base),
			ChangeFreq: sitemap.Weekly,
		},
	}

	for _, a := range articles {
		urlset.URLs = append(urlset.URLs, sitemap.URL{
			Loc:        fmt.Sprintf("%s/article/%s", configuration.Base, a.Slug),
			LastMod:    &a.UpdatedAt,
			ChangeFreq: sitemap.Monthly,
		})
	}

	raw, err := sitemap.Marshal(&urlset)
	if err != nil {
		panic(err)
	}

	w.Write(raw)
}
