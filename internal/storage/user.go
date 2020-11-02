package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

func (store *Store) GetUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error) {
	ctx, cancelCtx := setupContext(ctx)
	dbAccess, commitFunc, err := store.GetDAL(ctx)
	defer func() {
		if err := commitFunc(); err != nil && errVal == nil {
			errVal = err
		}
	}()
	if err != nil {
		cancelCtx()
		return nil, err
	}
	user, err = dbAccess.SelectUserByEmail(email)
	if err != nil {
		cancelCtx()
		return nil, err
	}
	return user, nil
}

// RegisterUser first inserts the user into the database, then queries the db and returns a UserProfile model
func (store *Store) RegisterUser(ctx context.Context, userIn *pkg.UserIn) (user *models.UserProfile, errVal error) {
	ctx, cancelCtx := setupContext(ctx)
	dbAccess, commitFunc, err := store.GetDAL(ctx)
	if err != nil {
		cancelCtx()
		return nil, err
	}
	if err = dbAccess.InsertUser(userIn); err != nil {
		cancelCtx()
		return nil, err
	}
	user, err = dbAccess.SelectUserByEmail(userIn.Email)
	if err != nil {
		return nil, err
	}
	if err = commitFunc(); err != nil {
		cancelCtx()
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user by his/her pub_id
func (store *Store) DeleteUser(ctx context.Context, userPubId *uuid.UUID) error {
	ctx, cancelCtx := setupContext(ctx)
	dbAccess, commitFunc, err := store.GetDAL(ctx)
	if err != nil {
		cancelCtx()
		return err
	}
	if err = dbAccess.DeleteUserByPubId(userPubId); err != nil {
		cancelCtx()
		return err
	}
	if err = commitFunc(); err != nil {
		cancelCtx()
		return err
	}
	return nil
}

func (store *Store) PromoteToLeader(ctx context.Context, userPubId *uuid.UUID, churchPubId *uuid.UUID) error {
	ctx, cancelCtx := setupContext(ctx)
	dbAccess, commitFunc, err := store.GetDAL(ctx)
	if err != nil {
		cancelCtx()
		return err
	}
	if err = dbAccess.MakeLeader(churchPubId, userPubId); err != nil {
		cancelCtx()
		return err
	}
	if err = commitFunc(); err != nil {
		cancelCtx()
		return err
	}
	return nil
}
