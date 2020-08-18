package session

import (
	"github.com/google/uuid"
)

type UserSession struct {
	UserUUID   *uuid.UUID `json:"userUuid"`
	ChurchUUID *uuid.UUID `json:"churchUuid"`
	Role       string    `json:"role"`
}


func NewSession() (userSession *UserSession){
	 return &UserSession{
		 UserUUID:   nil,
		 ChurchUUID: nil,
		 Role:       "",
	 }
}