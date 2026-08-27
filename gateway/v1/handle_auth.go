package v1

import (
	"encoding/json"
	"net/http"

	"omegaui.io/edgeflow/entities"
)

// Keeping it very simple, as my goal is to acheive
// distributed system traffic simulation.
// Open by default, generates access token for any user
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var user entities.UserEntity
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
