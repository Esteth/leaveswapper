package app

import (
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/chi"
	"github.com/unrolled/render"
)

type order struct {
	UserID string `db:"user_id"`
	Date   time.Time
	Type   string
}

type ordersResource struct {
	rendr *render.Render
	db    *sqlx.DB
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
	orders := []order{}
	err := rs.db.Select(&orders,
		rs.db.Rebind(
			`SELECT *
			 FROM orders
			 WHERE type = 'sell' 
         AND user_id = ?`),
		"adam.copp@gmail.com")
	if err != nil {
		panic(err)
	}
	rs.rendr.JSON(w, http.StatusOK, orders)
}

func (rs ordersResource) get(w http.ResponseWriter, r *http.Request) {
	order := order{}
	err := rs.db.Get(&order,
		rs.db.Rebind(
			`SELECT *
			 FROM orders
			 WHERE type = 'sell' 
         AND user_id = ?
			 LIMIT 1`),
		"adam.copp@gmail.com")
	if err != nil {
		panic(err)
	}
	rs.rendr.JSON(w, http.StatusOK, order)
}

func (rs ordersResource) create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order create"))
}

func (rs ordersResource) update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order update"))
}

func (rs ordersResource) delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order delete"))
}
