package endpoints

import (
	"encoding/json"
	"github.com/google/logger"
	"github.com/jsfan/hello-neighbour-api/pkg"
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
