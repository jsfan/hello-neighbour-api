package utils

import (
	"fmt"
	"net/http"
	"github.com/google/logger"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/session"
	"github.com/jsfan/hello-neighbour/internal/storage"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/internal/utils/crypto"
)

func userSessionFromProfile(profile *models.UserProfile) (userSession *session.UserSession) {
	var userUUID, churchUUID uuid.UUID
	var err error
	userUUID, err = uuid.Parse(profile.PubId)
	if err != nil {
		panic(fmt.Sprintf("Church's UUID from database could not be parsed: %+v", err))
	}
	churchUUID, err = uuid.Parse(profile.ChurchUUID)
	if err != nil {
		panic(fmt.Sprintf("User's UUID from database could not be parsed: %+v", err))
	}
	return &session.UserSession{
		UserUUID:   &userUUID,
		ChurchUUID: &churchUUID,
		Role:       profile.Role,
	}
}

func CheckBasicAuth(r *http.Request) (userSession *session.UserSession, authFail bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return nil, false
	}
	if password == "" {
		return nil, false
	}
	store, err := storage.GetStore()
	if err != nil {
		return nil, false
	}
	userProfile, err := store.GetUserByEmail(r.Context(), username)
	if err != nil {
		// TODO: Send a 500 unless this was a "Not found"
		logger.Error(err)
		return nil, true
	}
	if ok := crypto.CheckPassword([]byte(userProfile.PasswordHash), []byte(password)); !ok {
		return nil, true
	}
	return userSessionFromProfile(userProfile), false
}
