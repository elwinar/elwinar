package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// QuoteListHandler display a quote.
func QuoteListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var quotes []Quote

	err := database.Select(&quotes, "SELECT * FROM quotes ORDER BY id")
	if err != nil {
		panic(err)
	}
	fmt.Println(quotes)

	render(w, r, "quote_list", map[string]interface{}{
		"Title":  "Quotes",
		"Quotes": quotes,
	})
}
