package models

import "time"

// User - user model
type User struct {
	ID        uint64
	Balance   float64
	CreatedAt time.Time
}
