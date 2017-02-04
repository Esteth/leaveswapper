package leaveswapper

import "time"

// User entry in the datastore.
type User struct {
	Email   string
	Selling []time.Time
	Buying  []time.Time
}
