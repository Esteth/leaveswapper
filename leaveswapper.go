package leaveswapper

import "net/http"
import "html/template"
import "time"

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
	if err := rootTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<html>
  <head>
    <title>Rooty</title>
  </head>
  <body>
    <h1>Hiya, Go!</h1>
  </body>
</html>
`))
