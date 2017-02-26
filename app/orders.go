package app

import (
	"context"
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

	r.Route("/:date", func(r chi.Router) {
		r.Use(orderContext(rs.db))
		r.Get("/", rs.get)
		r.Put("/", rs.update)
		r.Delete("/", rs.delete)
	})

	return r
}

func orderContext(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			date := chi.URLParam(r, "date")
			order := order{}
			err := db.Get(&order,
				db.Rebind(
					`SELECT
					  *
					FROM
					  orders
					WHERE 
					  type = 'sell' 
						AND user_id = ?
						AND date = ?
					LIMIT 1`),
				"adam.copp@gmail.com",
				date)
			if err != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}
			ctx := context.WithValue(r.Context(), "order", order)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (rs ordersResource) list(w http.ResponseWriter, r *http.Request) {
	orders := []order{}
	err := rs.db.Select(&orders,
		rs.db.Rebind(
			`SELECT
			  *
			FROM
			  orders
			WHERE 
			  type = 'sell'
        AND user_id = ?`),
		"adam.copp@gmail.com")
	if err != nil {
		panic(err)
	}
	rs.rendr.JSON(w, http.StatusOK, orders)
}

func (rs ordersResource) get(w http.ResponseWriter, r *http.Request) {
	order := r.Context().Value("order")
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
