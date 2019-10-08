package db

import (
	"log"

	"github.com/andreylm/guru/pkg/models"
)

// Storage - db
type Storage interface {
	SaveBet(*models.Bet) uint64
	SaveDeposit(*models.Deposit) bool
	SaveTransaction(*models.Transaction) bool
	SaveUser(*models.User) bool
}

// NewMockStorage - creates new db connection
func NewMockStorage() Storage {
	return &MockRepository{}
}

// MockRepository - storage repository
type MockRepository struct {
	// db *sql.DB
}

// SaveBet - saving bet
func (s *MockRepository) SaveBet(model *models.Bet) uint64 {
	log.Println("DB save bet")
	return 1
}

// SaveDeposit - saving deposit
func (s *MockRepository) SaveDeposit(model *models.Deposit) bool {
	log.Println("DB save deposit")
	return true
}

// SaveTransaction - saving transaction
func (s *MockRepository) SaveTransaction(model *models.Transaction) bool {
	log.Println("DB save transaction")
	return true
}

// SaveUser - saving user
func (s *MockRepository) SaveUser(model *models.User) bool {
	log.Println("DB save user")
	return true
}
