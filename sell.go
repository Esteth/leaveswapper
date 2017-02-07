package leaveswapper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	appengineuser "google.golang.org/appengine/user"

	"google.golang.org/appengine"
)

type sale struct {
	time time.Time
}

func postNewSale(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if appengineuser.Current(ctx) == nil {
		url, err := appengineuser.LoginURL(ctx, "/")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		return
	}

	if r.Body == nil {
		http.Error(w, "Must send a request body", http.StatusBadRequest)
		return
	}

	var newSale sale
	err := json.NewDecoder(r.Body).Decode(&newSale)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Success")
}
