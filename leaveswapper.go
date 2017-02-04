package leaveswapper

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/save", save)
	http.HandleFunc("/sell", sell)
}

func root(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user User
	cursor := datastore.NewQuery("User").Run(ctx)
	_, err := cursor.Next(&user)
	if handleErr(err, w) {
		return
	}

	rootTemplate := template.Must(template.ParseFiles("templates/user.html"))
	err = rootTemplate.Execute(w, user)
	if handleErr(err, w) {
		return
	}
}
