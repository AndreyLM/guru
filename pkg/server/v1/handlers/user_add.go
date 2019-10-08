package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andreylm/guru/pkg/cache"
	"github.com/andreylm/guru/pkg/db"
	"github.com/andreylm/guru/pkg/errors"
	"github.com/andreylm/guru/pkg/models"
)

// CreateUserForm - create user form
type CreateUserForm struct {
	ID      uint64  `json:"id"`
	Balance float64 `json:"balance"`
	Token   string  `json:"token"`
}

// AddUser - creating new user
func AddUser(dbStorage db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form CreateUserForm
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

		user, err := getUserModel(&form)
		if err != nil {
			writeJSONResponse(w, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := cache.Storage.AddUser(user); err != nil {
			errors.DebugPrintf(err)
			writeJSONResponse(w, map[string]interface{}{"error": err.Error()})
			return
		}

		// Some logic saving user
		dbStorage.SaveUser(user)

		writeJSONResponse(w, map[string]interface{}{"error": ""})
	}
}

func getUserModel(form *CreateUserForm) (*models.User, error) {
	// Form validation logic
	if form.ID == 0 {
		return nil, errors.InvalidDataError
	}

	return &models.User{
		ID:      form.ID,
		Balance: form.Balance,
	}, nil
}
