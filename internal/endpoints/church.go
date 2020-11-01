package endpoints

import (
	"net/http"
)

// NewChurchRequest creates a new church
func NewChurchRequest(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement me
	w.WriteHeader(http.StatusCreated)
}
