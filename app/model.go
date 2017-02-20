package app

import "time"

// User entry in the datastore.
type user struct {
	Email   string
	Selling []time.Time
	Buying  []time.Time
}
