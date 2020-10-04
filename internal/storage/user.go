package storage

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
	"log"
)

func setupContext(ctx context.Context) (ctext context.Context, cancelCtx context.CancelFunc) {
	ctext, cancelCtx = context.WithCancel(ctx)
	return ctext, cancelCtx
}

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
		log.Printf("[ERROR] Database error: +%v", err)
		cancelCtx()
		return nil, err
	}
	return user, nil
}

// UserRegister will first insert the user into the database, then query the db and return a UserProfile model
func (store *Store) UserRegister(ctx context.Context, userIn *pkg.UserIn) (user *models.UserProfile, errVal error) {
	ctx, cancelCtx := setupContext(ctx)
	dbAccess, commitFunc, err := store.GetDAL(ctx)
	if err != nil {
		cancelCtx()
		return nil, err
	}
	if err = dbAccess.RegisterUser(userIn); err != nil {
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
