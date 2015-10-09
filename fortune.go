package main

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Fortune struct {
	ID     int64  `db:"id"`
	Text   string `db:"text"`
	Author string `db:"author"`
}

func FortuneHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var fortune Fortune

	err := database.Get(&fortune, "SELECT * FROM fortunes WHERE id >= (select ABS(RANDOM()) % MAX(id) + 1 FROM fortunes) LIMIT 1")
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	render(w, r, "fortune", map[string]interface{}{
		"Title":   "Fortune",
		"Fortune": fortune,
	})
}
