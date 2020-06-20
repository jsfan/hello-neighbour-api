package endpoints

import (
	"encoding/json"
	"github.com/jsfan/hello-neighbour/internal/config"
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

func sendErrorResponsee(w http.ResponseWriter, code int32, message string) {
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
		sendErrorResponsee(w, http.StatusBadRequest, "Access denied")
		return
	}
	if password == "" {
		sendErrorResponsee(w, http.StatusBadRequest, "Password cannot be empty")
		return
	}
	db := r.Context().Value(config.DatabaseConnection).(*storage.DBConnection)
	userProfile, err := db.GetUserByEmail(r.Context(), username)
	if err != nil {
		// todo: Send a 500 unless this was a "Not found"
		w.WriteHeader(http.StatusUnauthorized)
		sendErrorResponsee(w, http.StatusBadRequest, "Username or password incorrect")
		return
	}
	if ok := utils.CheckPassword([]byte(userProfile.PasswordHash), []byte(password)); !ok {
		sendErrorResponsee(w, http.StatusBadRequest, "Username or password incorrect")
		return
	}
	w.WriteHeader(http.StatusOK)
	successResp := pkg.Jwt{
		Jwt: "todo: implement this",
	}
	sendJsonResponse(w, successResp)
}
