package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/google/logger"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/pkg"
)

// NewChurchRequest allows users without an existing church affiliation to add a new church and become the leader of it
func NewChurchRequest(w http.ResponseWriter, r *http.Request) {
	userSession := r.Context().Value(config.SessionKey).(*session.UserSession)
	if (userSession.ChurchUUID != nil) {
		SendErrorResponse(w, http.StatusBadRequest, "You cannot request a new church if you currently belong to one.")
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Could not read request body: %+v", err)
		SendErrorResponse(w, http.StatusBadRequest, "")
		return
	}

	var churchIn *pkg.ChurchIn
	if err = json.Unmarshal(b, &churchIn); err != nil {
		logger.Errorf("Problem with unmarshaling JSON: %+v", err)
		SendErrorResponse(w, http.StatusBadRequest, "")
		return
	}

	db, err := storage.GetStore()
	if err != nil {
		logger.Errorf("Could not get db connection: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}

	church, err := db.AddChurch(r.Context(), churchIn)
	if err != nil {
		logger.Errorf("Database error: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	logger.Infof("Church request sent: %+v", church.PubId)
	w.WriteHeader(http.StatusCreated)
	SendJSONResponse(w, church)
}
