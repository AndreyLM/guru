package models

import (
	"time"
)

// BetType - transation type
type BetType string

const (
	// BetTypeWin - bet result
	BetTypeWin BetType = "Win"
	// BetTypeBet - making bet
	BetTypeBet BetType = "Bet"
)

// Bet - bet
type Bet struct {
	ID        uint64
	Type      BetType
	UserID    uint64
	Amount    float64
	CreatedAt time.Time
}
