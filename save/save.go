package save

import (
	"context"
	"net/http"
	"time"

	"github.com/esteth/leaveswapper/model"

	"github.com/go-martini/martini"
	"google.golang.org/appengine/datastore"
)

func Init(m martini.Router) {
	m.Get("/save", save)
}

func save(ctx context.Context) (int, string) {
	user := model.User{
		Email: "adam.copp@gmail.com",
		Selling: []time.Time{
			time.Date(2017, 8, 12, 0, 0, 0, 0, time.UTC),
		},
		Buying: []time.Time{
			time.Date(2017, 11, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "User", nil), &user)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	return http.StatusOK, "Saved new user!"
}
