package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"  
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

	churchPubId, err := uuid.Parse(church.PubId)
	if err != nil {
		// this shouldn't happen
		logger.Errorf("Cannot parse pub ID of church %s: %+v", church.PubId, err)
	}
	if err = db.PromoteToLeader(r.Context(), userSession.UserUUID, &churchPubId); err != nil {
		logger.Errorf("Problem promoting user %s to leader of church %s: %+v.", userSession.UserUUID.String(), churchPubId.String(), err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	logger.Infof("Church request sent: %+v", church.PubId)
	w.WriteHeader(http.StatusCreated)
	SendJSONResponse(w, church)
}

func ActivateChurch(w http.ResponseWriter, r *http.Request) {
	userSession := r.Context().Value(config.SessionKey).(*session.UserSession)
	if (userSession.Role != "admin") {
		SendErrorResponse(w, http.StatusForbidden, "You cannot change a church's activation status.")
	}

	db, err := storage.GetStore()
	if err != nil {
		logger.Errorf("Could not get db connection: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}

	churchUUIDStr := mux.Vars(r)["churchUuid"]
	churchUUID, err := uuid.Parse(churchUUIDStr)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid church UUID.")
		return
	}

	isActive, err := strconv.ParseBool(mux.Vars(r)["isActive"])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid active status.")
		return
	}

	if err = db.ChurchActivation(r.Context(), &churchUUID, isActive); err != nil {
		logger.Errorf("Could not change activation status of church %s: %+v", churchUUID.String(), err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	logger.Infof("Church %s has their activation status updated.", churchUUID.String())
	w.WriteHeader(http.StatusNoContent)
}