package sell

import (
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/sell", sell)
	http.HandleFunc("/sell/:date", sellDate)
}

func sell(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getSales(w, r)
	} else if r.Method == "POST" {
		postNewSale(w, r)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func sellDate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getSale(w, r)
	} else if r.Method == "DELETE" {
		deleteSale(w, r)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func getSales(w http.ResponseWriter, r *http.Request) {

}

func getSale(w http.ResponseWriter, r *http.Request) {

}

func postNewSale(w http.ResponseWriter, r *http.Request) {

}

func deleteSale(w http.ResponseWriter, r *http.Request) {

}
