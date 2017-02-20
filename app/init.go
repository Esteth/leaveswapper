package app

import (
	"net/http"

	"fmt"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/unrolled/render"
)

// Serves the api on the given address.
func ListenAndServe(addr string) {
	r := chi.NewRouter()
	rendr := render.New()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/orders", ordersResource{
		rendr: rendr,
	}.routes())

	fmt.Printf("Now listening on %s\n", addr)
	http.ListenAndServe(addr, r)
}
