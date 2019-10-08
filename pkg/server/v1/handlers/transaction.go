package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/andreylm/guru/pkg/cache"
	"github.com/andreylm/guru/pkg/db"
	"github.com/andreylm/guru/pkg/errors"
	"github.com/andreylm/guru/pkg/models"
)

// TransactionRequest - transaction request
type TransactionRequest struct {
	ID     uint64         `json:"transactionId"`
	UserID uint64         `json:"userId"`
	Amount float64        `json:"amount"`
	Type   models.BetType `json:"type"`
	Token  string         `json:"token"`
}

// Transaction - transaction
func Transaction(dbStorage db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mu sync.Mutex
		var request TransactionRequest
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&request); err != nil {
			errors.DebugPrintf(err)
			writeJSONResponse(w, map[string]interface{}{"error": errors.InternalServerError.Error()})
			return
		}

		if !validateToken(request.Token) {
			writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidTokenError.Error()})
			return
		}

		if !validateTransaction(&request) {
			writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidDataError.Error()})
			return
		}

		user, err := cache.Storage.GetUser(request.UserID)
		if err != nil {
			writeJSONResponse(w, map[string]interface{}{"error": err.Error()})
			return
		}

		if request.Type == models.BetTypeBet && request.Amount > user.Balance {
			writeJSONResponse(w, map[string]interface{}{"error": "Your balance is too small for making such bet"})
			return
		}

		betID, err := saveBet(dbStorage, &request)
		if err != nil {
			errors.DebugPrintf(err)
			writeJSONResponse(w, map[string]interface{}{"error": err.Error()})
			return
		}
		saveTransaction(dbStorage, betID, user.Balance, &request)

		mu.Lock()
		defer mu.Unlock()
		if request.Type == models.BetTypeBet {
			user.Balance -= request.Amount
			if err := cache.Storage.UpdateUserStats(user.ID, request.Amount, cache.UserChangesBet); err != nil {
				errors.DebugPrintf(err)
			}
		} else {
			user.Balance += request.Amount
			if err := cache.Storage.UpdateUserStats(user.ID, request.Amount, cache.UserChangesWin); err != nil {
				errors.DebugPrintf(err)
			}
		}

		cache.Storage.AddModifiedUser(user)

		writeJSONResponse(w, map[string]interface{}{
			"error":   "",
			"balance": user.Balance,
		})
	}
}

func validateTransaction(model *TransactionRequest) bool {
	return model.ID != 0 && model.UserID != 0 && model.Amount != 0 && (model.Type == models.BetTypeBet || model.Type == models.BetTypeWin)
}

func saveBet(dbStorage db.Storage, model *TransactionRequest) (uint64, error) {
	bet := &models.Bet{
		UserID:    model.UserID,
		Type:      model.Type,
		Amount:    model.Amount,
		CreatedAt: time.Now(),
	}
	return dbStorage.SaveBet(bet), nil
}

func saveTransaction(dbStorage db.Storage, betID uint64, userBalance float64, model *TransactionRequest) bool {
	transaction := &models.Transaction{
		UserID:        model.UserID,
		BetID:         betID,
		Amount:        model.Amount,
		BalanceBefore: userBalance,
		BalanceAfter:  userBalance + model.Amount,
		CreatedAt:     time.Now(),
	}
	return dbStorage.SaveTransaction(transaction)
}
