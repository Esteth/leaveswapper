package sell

import "net/http"

func Sell(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		get(w, r)
	} else if r.Method == "POST" {
		post(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {

}

func post(w http.ResponseWriter, r *http.Request) {

}
