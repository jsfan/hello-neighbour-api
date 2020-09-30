package endpoints

import (
	"encoding/json"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/pkg"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/jsfan/hello-neighbour/internal/storage"
)

func SendJsonResponse(w http.ResponseWriter, jsonIn interface{}) {
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
		log.Printf("[ERROR] Could not build JWT: %+v", err)
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
		log.Printf("[ERROR] Could not read request body: %+v", err)
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var userIn *pkg.UserIn
	if err = json.Unmarshal(b, &userIn); err != nil {
		log.Printf("[ERROR] Problem with unmarshaling JSON: %+v", err)
		SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	db, err := storage.GetConnection()
	if err != nil {
		log.Printf("[ERROR] Could not get db connection: %+v", err.Error())
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := db.UserRegister(r.Context(), userIn)
	if err != nil {
		log.Printf("[ERROR] Database error: %+v", err.Error())
		SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(201)
	SendJsonResponse(w, user)
}