package leaveswapper

import (
	"encoding/json"
	"net/http"
	"time"

	appengineuser "google.golang.org/appengine/user"

	"golang.org/x/net/context"
)

type sale struct {
	time time.Time
}

func getSales() {

}

func getSale() {

}

func postNewSale(ctx context.Context, w http.ResponseWriter, r *http.Request) (int, string) {
	if appengineuser.Current(ctx) == nil {
		url, err := appengineuser.LoginURL(ctx, "/")
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		return http.StatusTemporaryRedirect, ""
	}

	if r.Body == nil {
		return http.StatusBadRequest, "Must send a request body"
	}

	var newSale sale
	err := json.NewDecoder(r.Body).Decode(&newSale)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return 200, "Success"
}

func deleteSale(w http.ResponseWriter, r *http.Request) {

}
