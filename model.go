package leaveswapper

import (
	"errors"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	appengineuser "google.golang.org/appengine/user"
)

// User entry in the datastore.
type user struct {
	Email   string
	Selling []time.Time
	Buying  []time.Time
}

func getUser(ctx context.Context) (user, error) {
	var u user

	if appengineuser.Current(ctx) == nil {
		return u, errors.New("Must be signed in")
	}

	email := appengineuser.Current(ctx).Email
	key := datastore.NewKey(ctx, "User", email, 0, nil)
	err := datastore.Get(ctx, key, &u)
	if err == datastore.ErrNoSuchEntity {
		u.Email = email
		datastore.Put(ctx, key, &u)
	}
	return u, nil
}
