package leaveswapper

import (
	"net/http"
	"time"

	"fmt"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func save(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	user := user{
		Email: "adam.copp@gmail.com",
		Selling: []time.Time{
			time.Date(2017, 8, 12, 0, 0, 0, 0, time.UTC),
		},
		Buying: []time.Time{
			time.Date(2017, 11, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "User", nil), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Saved new user!")
}
