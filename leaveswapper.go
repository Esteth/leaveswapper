package leaveswapper

import (
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// User entry in the datastore.
type User struct {
	Email   string
	Selling []time.Time
	Buying  []time.Time
}

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	user := User{
		Email: "adam.copp@gmail.com",
		Selling: []time.Time{
			time.Date(2017, 8, 12, 0, 0, 0, 0, time.UTC),
		},
		Buying: []time.Time{
			time.Date(2017, 11, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "User", nil), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var retrievedUser User
	if err = datastore.Get(ctx, key, &retrievedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := rootTemplate.Execute(w, retrievedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<html>
  <head>
    <title>Rooty</title>
  </head>
  <body>
    <h1>{{.Email}}</h1>
		<h2>Selling</h2>
		<ul>
			{{range $sell := .Selling}}
				{{$sell}}
			{{end}}
		</ul>
		<h2>Buying</h2>
		<ul>
			{{range $buy := .Buying}}
				{{$buy}}
			{{end}}
		</ul>
  </body>
</html>
`))
