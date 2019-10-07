package models

// TransactionType - transation type
type TransactionType string

var (
	// TransactionTypeWin - bet result
	TransactionTypeWin TransactionType = "Win"
	// TransactionTypeBet - making bet
	TransactionTypeBet TransactionType = "Bet"
)

// Transaction - transaction
type Transaction struct {
	ID     int
	UserID int
	Type   TransactionType
	Amout  float64
	Token  string
}
