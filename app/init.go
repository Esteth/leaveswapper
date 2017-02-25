package app

import (
	"net/http"

	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/unrolled/render"

	_ "github.com/lib/pq"
)

// Serves the api on the given address.
func ListenAndServe(addr string, dburl string) error {
	r := chi.NewRouter()
	rendr := render.New()

	db, err := sqlx.Open("postgres", dburl)
	if err != nil {
		return err
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/orders", ordersResource{
		rendr: rendr,
		db:    db,
	}.routes())

	fmt.Printf("Now listening on %s\n", addr)
	http.ListenAndServe(addr, r)

	return nil
}
