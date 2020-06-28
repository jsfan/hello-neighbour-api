package session

import "github.com/google/uuid"

type UserSession struct {
	UserUUID   uuid.UUID `json:"userUuid"`
	ChurchUUID uuid.UUID `json:"churchUuid"`
}
