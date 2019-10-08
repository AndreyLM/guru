package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andreylm/guru/pkg/cache"
	"github.com/andreylm/guru/pkg/db"
	"github.com/andreylm/guru/pkg/errors"
)

// GetUserForm - get user form
type GetUserForm struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

// GetUser - creating new user
func GetUser(dbStorage db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form GetUserForm
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&form); err != nil {
			errors.DebugPrintf(err)
			writeJSONResponse(w, map[string]interface{}{"error": errors.InternalServerError.Error()})
			return
		}

		if !validateToken(form.Token) {
			writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidTokenError.Error()})
			return
		}

		if !validateGetUserForm(&form) {
			writeJSONResponse(w, map[string]interface{}{"error": errors.InvalidDataError.Error()})
			return
		}

		info, err := getUserInfo(&form)
		if err != nil {
			writeJSONResponse(w, map[string]interface{}{"error": err.Error()})
			return
		}

		writeJSONResponse(w, info)
	}
}

func validateGetUserForm(form *GetUserForm) bool {
	return form.ID != 0
}

func getUserInfo(form *GetUserForm) (map[string]interface{}, error) {
	user, err := cache.Storage.GetUser(form.ID)
	if err != nil {
		return nil, err
	}
	userStatistics := cache.Storage.GetUserStatistics(form.ID)

	return map[string]interface{}{
		"id":           form.ID,
		"balance":      user.Balance,
		"depositCount": userStatistics.DepositCount,
		"depositSum":   userStatistics.DepositSum,
		"betCount":     userStatistics.BetCount,
		"betSum":       userStatistics.BetSum,
		"winCount":     userStatistics.WinCount,
		"winSum":       userStatistics.WinSum,
	}, nil
}
