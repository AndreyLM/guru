package models

// Statistics - statistics model
type Statistics struct {
	DepositCount uint    `json:"depositCount"`
	DepositSum   float64 `json:"depositSum"`
	BetCount     uint    `json:"betCount"`
	BetSum       float64 `json:"betSum"`
	WinCount     uint    `json:"winCount"`
	WinSum       float64 `json:"winSum"`
}
