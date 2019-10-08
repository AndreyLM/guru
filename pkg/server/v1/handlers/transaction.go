package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/andreylm/guru/pkg/cache"
	"github.com/andreylm/guru/pkg/errors"
	"github.com/andreylm/guru/pkg/models"
)

// Transaction - transaction
func Transaction(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex
	var transaction models.Transaction
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&transaction); err != nil {
		errors.DebugPrintf(err)
		writeJSONResponse(w, map[string]interface{}{"error": errors.InternalServerError.Error()})
		return
	}

	if !validateToken(transaction.Token) {
		writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidTokenError.Error()})
		return
	}

	if !validateTransaction(&transaction) {
		writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidDataError.Error()})
		return
	}

	user, err := cache.Storage.GetUser(transaction.UserID)
	if err != nil {
		writeJSONResponse(w, map[string]interface{}{"error": err.Error()})
		return
	}

	if transaction.Type == models.TransactionTypeBet && transaction.Amount > user.Balance {
		writeJSONResponse(w, map[string]interface{}{"error": "Your balance is too small for making such bet"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if transaction.Type == models.TransactionTypeBet {
		user.Balance -= transaction.Amount
		if err := cache.Storage.UpdateUserStats(user.ID, transaction.Amount, cache.UserChangesBet); err != nil {
			errors.DebugPrintf(err)
		}
	} else {
		user.Balance += transaction.Amount
		if err := cache.Storage.UpdateUserStats(user.ID, transaction.Amount, cache.UserChangesWin); err != nil {
			errors.DebugPrintf(err)
		}
	}

	cache.Storage.AddModifiedUser(user)

	writeJSONResponse(w, map[string]interface{}{
		"error":   "",
		"balance": user.Balance,
	})
}

func validateTransaction(model *models.Transaction) bool {
	return model.ID != 0 && model.UserID != 0 && model.Amount != 0 && (model.Type == models.TransactionTypeWin || model.Type == models.TransactionTypeBet)
}
