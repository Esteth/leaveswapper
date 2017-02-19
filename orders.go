package leaveswapper

import "github.com/pressly/chi"
import "net/http"

type ordersResource struct{}

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
	w.Write([]byte("List of orders"))
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
