package leaveswapper

import (
	"html/template"
	"net/http"

	"github.com/esteth/leaveswapper/model"
	"github.com/esteth/leaveswapper/save"
	"github.com/esteth/leaveswapper/sell"
	"github.com/esteth/leaveswapper/utils"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// Run is the main handler function.
func Run() {
	http.HandleFunc("/", root)
	save.RegisterHandlers()
	sell.RegisterHandlers()
}

func root(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user model.User
	cursor := datastore.NewQuery("User").Run(ctx)
	_, err := cursor.Next(&user)
	if utils.HandleErr(err, w) {
		return
	}

	rootTemplate := template.Must(template.ParseFiles("templates/user.html"))
	err = rootTemplate.Execute(w, user)
	if utils.HandleErr(err, w) {
		return
	}
}
