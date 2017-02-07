package leaveswapper

import (
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// Init is the main handler function.
func Init(m martini.Router) {
	m.Get("/", root)
	m.Get("/save", save)
	m.Get("/sell", getSales)
	m.Post("/sell", postNewSale)
	m.Get("/sell/:date", getSale)
}

func root(w http.ResponseWriter, r *http.Request) (int, string) {
	ctx := appengine.NewContext(r)

	var user user
	cursor := datastore.NewQuery("User").Run(ctx)
	_, err := cursor.Next(&user)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	rootTemplate := template.Must(template.ParseFiles("templates/user.html"))
	err = rootTemplate.Execute(w, user)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, ""
}
