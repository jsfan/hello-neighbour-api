package session

import (
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/config"
	"net/http"
)

type UserSession struct {
	UserUUID   uuid.UUID `json:"userUuid"`
	ChurchUUID uuid.UUID `json:"churchUuid"`
	Role       string    `json:"role"`
}


func GetSession(r *http.Request) (userSession *UserSession){
	 return r.Context().Value(config.SessionKey).(*UserSession)
}