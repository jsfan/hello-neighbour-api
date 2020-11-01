package endpoints

import (
	"encoding/json"
	"github.com/google/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/pkg"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	userSession := r.Context().Value(config.SessionKey).(*session.UserSession)
	jwtRef := session.NewJWT()
	err := jwtRef.Build(userSession)
	if err != nil {
		logger.Errorf("Could not build JWT: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
	}
	successResp := pkg.Jwt{
		Jwt: jwtRef.GetRaw(),
	}
	SendJSONResponse(w, successResp)
}

// DefaultUserRegister is the default signup when creating a new church
func DefaultUserRegister(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Could not read request body: %+v", err)
		SendErrorResponse(w, http.StatusBadRequest, "")
		return
	}

	var userIn *pkg.UserIn
	if err = json.Unmarshal(b, &userIn); err != nil {
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

	user, err := db.RegisterUser(r.Context(), userIn)
	if err != nil {
		logger.Errorf("Database error: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	logger.Infof("User created: %+v", user.PubId)
	w.WriteHeader(http.StatusCreated)
	SendJSONResponse(w, user)
}

// DeleteUserAccount deletes a user and all his/her assets
func DeleteUserAccount(w http.ResponseWriter, r *http.Request) {
	userUUIDdStr := mux.Vars(r)["userUuid"]
	userUUID, err := uuid.Parse(userUUIDdStr)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid user UUID.")
		return
	}
	userSession := r.Context().Value(config.SessionKey).(*session.UserSession)
	currentUserUUID := userSession.UserUUID
	if (&userUUID != currentUserUUID) && (userSession.Role != "admin") {
		SendErrorResponse(w, http.StatusForbidden, "You cannot delete that user.")
	}
	db, err := storage.GetStore()
	if err != nil {
		logger.Errorf("Could not get db connection: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}

	if err = db.DeleteUser(r.Context(), &userUUID); err != nil {
		logger.Errorf("Could not delete user %s: +%v", currentUserUUID.String(), err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	logger.Infof("User %s was deleted.", currentUserUUID.String())
	w.WriteHeader(http.StatusNoContent)
}
