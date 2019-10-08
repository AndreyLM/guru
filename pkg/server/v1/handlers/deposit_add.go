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

// AddDeposit - add deposit
func AddDeposit(dbStorage db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mu sync.Mutex
		var deposit models.Deposit
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&deposit); err != nil {
			errors.DebugPrintf(err)
			writeJSONResponse(w, map[string]interface{}{"error": errors.InternalServerError.Error()})
			return
		}

		if !validateToken(deposit.Token) {
			writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidTokenError.Error()})
			return
		}

		if !validateDeposit(&deposit) {
			writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidDataError.Error()})
			return
		}

		user, err := cache.Storage.GetUser(deposit.UserID)
		if err != nil {
			writeJSONResponse(w, map[string]interface{}{"error": err.Error()})
			return
		}

		deposit.CreatedAt = time.Now()
		deposit.BalanceBefore = user.Balance
		deposit.BalanceAfter = user.Balance + deposit.Amount
		dbStorage.SaveDeposit(&deposit)

		mu.Lock()
		defer mu.Unlock()
		user.Balance += deposit.Amount

		cache.Storage.AddModifiedUser(user)
		if err := cache.Storage.UpdateUserStats(user.ID, deposit.Amount, cache.UserChangesDeposit); err != nil {
			errors.DebugPrintf(err)
		}

		writeJSONResponse(w, map[string]interface{}{
			"error":   "",
			"balance": user.Balance,
		})
	}
}

func validateDeposit(model *models.Deposit) bool {
	return model.UserID != 0 && model.ID != 0 && model.Amount != 0
}
