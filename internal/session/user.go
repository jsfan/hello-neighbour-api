package session

import (
	"github.com/jsfan/hello-neighbour-api/internal/config"
)

func NewSession() (userSession *config.UserSession) {
	return &config.UserSession{
		UserUUID:   nil,
		ChurchUUID: nil,
		Role:       "",
	}
}
