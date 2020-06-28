package endpoints

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/internal/utils"
	"github.com/jsfan/hello-neighbour/pkg"
	"log"
	"net/http"
)

func sendJsonResponse(w http.ResponseWriter, jsonIn interface{}) {
	if jsonIn != nil {
		resp, err := json.Marshal(&jsonIn)
		if err != nil {
			log.Printf("[ERROR] Could not marshal JSON response: %+v", err)
		}
		if _, err := w.Write(resp); err != nil {
			log.Printf("[ERROR] Could not send JSON response: %+v", err)
		}
	}
}

func sendErrorResponse(w http.ResponseWriter, code int32, message string) {
	w.WriteHeader(int(code))
	errResp := pkg.ErrorResponse{
		Code:    code,
		Message: message,
	}
	sendJsonResponse(w, errResp)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		sendErrorResponse(w, http.StatusBadRequest, "Access denied")
		return
	}
	if password == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Password cannot be empty")
		return
	}
	db := r.Context().Value(config.DatabaseConnection).(*storage.DBConnection)
	userProfile, err := db.GetUserByEmail(r.Context(), username)
	if err != nil {
		// todo: Send a 500 unless this was a "Not found"
		w.WriteHeader(http.StatusUnauthorized)
		sendErrorResponse(w, http.StatusBadRequest, "Username or password incorrect")
		return
	}
	if ok := utils.CheckPassword([]byte(userProfile.PasswordHash), []byte(password)); !ok {
		sendErrorResponse(w, http.StatusBadRequest, "Username or password incorrect")
		return
	}
	userUUID, err := uuid.Parse(userProfile.PubId)
	if err != nil {
		log.Printf("Could not parse user's UUID: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		sendErrorResponse(w, http.StatusInternalServerError, "Database error")
	}
	churchUUID, err := uuid.Parse(userProfile.ChurchUUID)
	if err != nil {
		log.Printf("Could not parse user's UUID: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		sendErrorResponse(w, http.StatusInternalServerError, "Database error")
	}
	w.WriteHeader(http.StatusOK)
	jwtRef := session.GetJWTWrapper()
	err = jwtRef.Build(&session.UserSession{
		UserUUID:   userUUID,
		ChurchUUID: churchUUID,
	})
	if err != nil {
		log.Printf("Could not build JWT: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
	}
	successResp := pkg.Jwt{
		Jwt: jwtRef.GetRaw(),
	}
	sendJsonResponse(w, successResp)
}
