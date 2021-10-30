package endpoints

import (
	"encoding/json"
	"github.com/jsfan/hello-neighbour-api/internal/interfaces"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

// NewChurchRequest allows users without an existing church affiliation to add a new church and become the leader of it
func NewChurchRequest(w http.ResponseWriter, r *http.Request) {
	userSession := r.Context().Value(config.SessionKey).(*config.UserSession)
	if userSession.ChurchUUID != nil {
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

	store := r.Context().Value(config.MasterStore).(interfaces.DataInterface)

	church, err := store.AddChurch(r.Context(), churchIn)
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
	if err = store.PromoteToLeader(r.Context(), userSession.UserUUID, &churchPubId); err != nil {
		logger.Errorf("Problem promoting user %s to leader of church %s: %+v.", userSession.UserUUID.String(), churchPubId.String(), err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	logger.Infof("Church request sent: %+v", church.PubId)
	w.WriteHeader(http.StatusCreated)
	SendJSONResponse(w, church)
}

func ActivateChurch(w http.ResponseWriter, r *http.Request) {
	userSession := r.Context().Value(config.SessionKey).(*config.UserSession)
	if userSession.Role != "admin" {
		SendErrorResponse(w, http.StatusForbidden, "You cannot change a church's activation status.")
		return
	}

	store := r.Context().Value(config.MasterStore).(interfaces.DataInterface)

	churchUUIDStr := mux.Vars(r)["churchUuid"]
	churchUUID, err := uuid.Parse(churchUUIDStr)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid church UUID.")
		return
	}

	isActive, err := strconv.ParseBool(r.URL.Query()["isActive"][0])
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid active status.")
		return
	}

	if err = store.ActivateChurch(r.Context(), &churchUUID, isActive); err != nil {
		logger.Errorf("Could not change activation status of church %s: %+v", churchUUID.String(), err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	logger.Infof("Church %s has their activation status updated.", churchUUID.String())
	w.WriteHeader(http.StatusOK)
}
