package models

import (
	"time"
)

// Transaction - transaction
type Transaction struct {
	ID            uint64    `json:"transactionId"`
	UserID        uint64    `json:"userId"`
	BetID         uint64    `json:"betID"`
	Amount        float64   `json:"amount"`
	BalanceBefore float64   `json:"balanceBefore"`
	BalanceAfter  float64   `json:"balanceAfter"`
	CreatedAt     time.Time `json:"createdAt"`
}
