package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/rest/model"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
)

func (store *Store) GetUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error) {
	user, err := store.DAL.SelectUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// RegisterUser first inserts the user into the database, then queries the db and returns a UserProfile model
func (store *Store) RegisterUser(ctx context.Context, userIn *model.UserIn) (user *models.UserProfile, errVal error) {
	if err := store.DAL.BeginTransaction(); err != nil {
		return nil, err
	}
	defer func() {
		if errVal != nil {
			store.DAL.CancelTransaction() // nolint:errcheck
			return
		}
		errVal = store.DAL.CompleteTransaction()
	}()

	if err := store.DAL.InsertUser(ctx, userIn); err != nil {
		return nil, err
	}
	user, err := store.DAL.SelectUserByEmail(ctx, userIn.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user by his/her pub_id
func (store *Store) DeleteUser(ctx context.Context, userPubId *uuid.UUID) error {
	return store.DAL.DeleteUserByPubId(ctx, userPubId)
}

// PromoteToLeader will promote a given member of a church to a leader role
func (store *Store) PromoteToLeader(ctx context.Context, userPubId *uuid.UUID, churchPubId *uuid.UUID) error {
	return store.DAL.MakeLeader(ctx, churchPubId, userPubId)
}
