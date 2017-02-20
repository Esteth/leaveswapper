package app

import "github.com/pressly/chi"
import "net/http"
import "time"

import "github.com/unrolled/render"

type order struct {
	Date time.Time
}

type ordersResource struct {
	rendr *render.Render
}

func (rs ordersResource) routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.list)
	r.Post("/", rs.create)
	r.Put("/", rs.delete)

	r.Route("/:id", func(r chi.Router) {
		r.Get("/", rs.get)
		r.Put("/", rs.update)
		r.Delete("/", rs.delete)
	})

	return r
}

func (rs ordersResource) list(w http.ResponseWriter, r *http.Request) {
	orders := []order{
		order{
			Date: time.Date(2017, 11, 25, 0, 0, 0, 0, time.UTC),
		}}
	rs.rendr.JSON(w, http.StatusOK, orders)
}

func (rs ordersResource) create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order create"))
}

func (rs ordersResource) delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order delete"))
}

func (rs ordersResource) get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order get"))
}

func (rs ordersResource) update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order update"))
}
