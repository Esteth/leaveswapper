package leaveswapper

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/save", save).Methods("GET")
	router.HandleFunc("/sell", postNewSale).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	http.Handle("/", n)
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
