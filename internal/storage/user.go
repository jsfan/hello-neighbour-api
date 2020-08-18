package storage

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"log"
)

func setupContextAndDbConnection(ctx context.Context) (ctext context.Context, cancelCtx context.CancelFunc, db *DBConnection) {
	ctext, cancelCtx = context.WithCancel(ctx)
	db, err := GetConnection()
	if err != nil { // panic here as we have checked before
		panic("Database connection unavailable")
	}
	return ctext, cancelCtx, db
}

func (conn *DBConnection) GetUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error) {
	ctx, cancelCtx, db := setupContextAndDbConnection(ctx)
	dbAccess, commitFunc, err := dal.GetDal(ctx, db.Db)
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
