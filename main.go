package leaveswapper

import (
	"html/template"
	"net/http"

	"github.com/pressly/chi"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	r := chi.NewRouter()

	r.Get("/", root)
	r.Get("/save", save)
	r.Post("/sell", postNewSale)

	http.Handle("/", r)
}

func root(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user user
	cursor := datastore.NewQuery("User").Run(ctx)
	_, err := cursor.Next(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rootTemplate := template.Must(template.ParseFiles("templates/user.html"))
	err = rootTemplate.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
