package models

import "time"

// User - user model
type User struct {
	ID        uint64    `json:"id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
