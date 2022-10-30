package utils

import (
	"context"
	"fmt"

	"github.com/golang/glog"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/config"
	"github.com/jsfan/hello-neighbour-api/internal/interfaces"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/internal/utils/crypto"
)

func userSessionFromProfile(profile *models.UserProfile) (userSession *config.UserSession) {
	var userUUID uuid.UUID
	var churchUUID *uuid.UUID
	var err error
	userUUID, err = uuid.Parse(profile.PubId)
	if err != nil {
		panic(fmt.Sprintf("User's UUID from database could not be parsed: %+v", err))
	}
	if profile.ChurchUUID != "" {
		tempUUID, err := uuid.Parse(profile.ChurchUUID)
		if err != nil {
			panic(fmt.Sprintf("Church's UUID from database could not be parsed: %+v", err))
		}
		churchUUID = &tempUUID
	}
	return &config.UserSession{
		UserUUID:   &userUUID,
		ChurchUUID: churchUUID,
		Role:       profile.Role,
	}
}

func CheckBasicAuth(ctx context.Context, store interfaces.DataInterface, username, password string) (*config.UserSession, bool) {
	if password == "" {
		return nil, false
	}

	userProfile, err := store.GetUserByEmail(ctx, username)
	if err != nil {
		// TODO: Send a 500 unless this was a "Not found"
		glog.Error(err)
		return nil, false
	}
	if ok := crypto.CheckPassword([]byte(userProfile.PasswordHash), []byte(password)); !ok {
		return nil, false
	}
	return userSessionFromProfile(userProfile), true
}
