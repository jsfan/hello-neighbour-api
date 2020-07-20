package endpoints

import (
	"encoding/json"
	"github.com/jsfan/hello-neighbour/internal/config"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/pkg"
	"log"
	"net/http"
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
		log.Printf("Could not build JWT: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
	}
	successResp := pkg.Jwt{
		Jwt: jwtRef.GetRaw(),
	}
	SendJsonResponse(w, successResp)
}
