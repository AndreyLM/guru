package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andreylm/guru/pkg/errors"
)

func writeJSONResponse(w http.ResponseWriter, data map[string]interface{}) {
	packet, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, errors.InternalServerError.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
