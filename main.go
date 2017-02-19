package leaveswapper

import (
	"net/http"

	"fmt"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

// Serves the api on the given address.
func ListenAndServe(addr string) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/orders", ordersResource{}.routes())

	fmt.Printf("Now listening on %s\n", addr)
	http.ListenAndServe(addr, r)
}
