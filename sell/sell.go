package sell

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

type sale struct {
	times []time.Time
}

// Init creates routes for adding sales.
func Init(m martini.Router) {
	m.Get("/sell", getSales)
	m.Post("/sell", postNewSale)

	m.Get("/sell/:date", getSale)
}

func getSales() {

}

func getSale() {

}

func postNewSale(r *http.Request) (int, string) {
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
