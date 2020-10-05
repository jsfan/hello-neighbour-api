package endpoints

import (
	"encoding/json"
	"github.com/google/logger"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/pkg"
	"io/ioutil"
	"net/http"
)

func SendJsonResponse(w http.ResponseWriter, jsonIn interface{}) {
	if jsonIn != nil {
		resp, err := json.Marshal(&jsonIn)
		if err != nil {
			logger.Errorf("Could not marshal JSON response: %+v", err)
		}
		if _, err := w.Write(resp); err != nil {
			logger.Errorf("Could not send JSON response: %+v", err)
		}
	}
}

func SendErrorResponse(w http.ResponseWriter, code int32, message string) {
	w.WriteHeader(int(code))
	errResp := pkg.ErrorResponse{
		Code:    code,
		Message: message,
	}
	SendJsonResponse(w, errResp)
}

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
	SendJsonResponse(w, successResp)
}

// DefaultUserRegister is the default signup when creating a new church
func DefaultUserRegister(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Could not read request body: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "")
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

	user, err := db.UserRegister(r.Context(), userIn)
	if err != nil {
		logger.Errorf("Database error: %+v", err)
		SendErrorResponse(w, http.StatusBadRequest, "")
		return
	}
	w.WriteHeader(201)
	SendJsonResponse(w, user)
}
