package models

// Deposit - deposit model
type Deposit struct {
	ID        int
	DepositID int
	UserID    int
	Amount    int
	Token     string
}
