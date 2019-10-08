package models

// Deposit - deposit model
type Deposit struct {
	ID     int     `json:"depositId"`
	UserID uint64  `json:"userId"`
	Amount float64 `json:"amount"`
	Token  string  `json:"token"`
}
