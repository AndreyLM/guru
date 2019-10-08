package models

import "time"

// Deposit - deposit model
type Deposit struct {
	ID            int       `json:"depositId"`
	UserID        uint64    `json:"userId"`
	Amount        float64   `json:"amount"`
	Token         string    `json:"token"`
	BalanceBefore float64   `json:"balanceBefore"`
	BalanceAfter  float64   `json:"balanceAfter"`
	CreatedAt     time.Time `json:"createdAt"`
}
