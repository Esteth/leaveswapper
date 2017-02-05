package sell

import (
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/sell", sell)
}

func sell(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		get(w, r)
	} else if r.Method == "POST" {
		post(w, r)
	} else if r.Method == "DELETE" {
		delete(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {

}

func post(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}
