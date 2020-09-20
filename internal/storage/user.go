package storage

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"log"
	"github.com/jsfan/hello-neighbour/pkg"
)

func setupContextAndDbConnection(ctx context.Context) (ctext context.Context, cancelCtx context.CancelFunc, store *Store) {
	ctext, cancelCtx = context.WithCancel(ctx)
	conn, err := GetStore()
	if err != nil { // panic here as we have checked before
		panic("Database connection unavailable")
	}
	return ctext, cancelCtx, conn
}

func (store *Store) GetUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error) {
	ctx, cancelCtx, store := setupContextAndDbConnection(ctx)
	dbAccess, commitFunc, err := dal.GetDAL(ctx)
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
func (conn *DBConnection) UserRegister(ctx context.Context, userIn *pkg.UserIn) (user *models.UserProfile, errVal error) {
	ctx, cancelCtx, db := setupContextAndDbConnection(ctx)
	dbAccess, commitFunc, err := dal.GetDal(ctx, db.Db)
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
