package models

// TransactionType - transation type
type TransactionType string

const (
	// TransactionTypeWin - bet result
	TransactionTypeWin TransactionType = "Win"
	// TransactionTypeBet - making bet
	TransactionTypeBet TransactionType = "Bet"
)

// Transaction - transaction
type Transaction struct {
	ID     uint64          `json:"transactionId"`
	UserID uint64          `json:"userId"`
	Type   TransactionType `json:"type"`
	Amount float64         `json:"amount"`
	Token  string          `json:"token"`
}
