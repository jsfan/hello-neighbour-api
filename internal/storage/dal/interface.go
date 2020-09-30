package dal

import (
	"context"
	"database/sql"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

type Dal struct {
	ctx context.Context
	tx  *sql.Tx
}

type DalInterface interface {
	SelectUserByEmail(email string) (user models.UserProfile, errVal error)
	RegisterUser(userIn *pkg.UserIn) (error)
}

func GetDal(ctx context.Context, db *sql.DB) (dbAccess *Dal, commit func() error, errVal error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	return &Dal{
		ctx: ctx,
		tx:  tx,
	}, func() error { return tx.Commit() }, nil
}
