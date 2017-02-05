package appengine

import (
	"net/http"

	"google.golang.org/appengine"

	"github.com/esteth/leaveswapper"
	"github.com/go-martini/martini"
)

func init() {
	m := martini.Classic()

	m.Use(appEngine)

	leaveswapper.Init(m)

	http.Handle("/", m)
}

func appEngine(c martini.Context, r *http.Request) {
	c.Map(appengine.NewContext(r))
}
