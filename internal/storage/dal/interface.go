package dal

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

type DAL struct {
	Db  *sql.DB
	ctx context.Context
	tx  *sql.Tx
}

type AccessInterface interface {
	SetupDal(ctx context.Context) (commit func() error, rollback func() error, errVal error)
	SelectUserByEmail(email string) (user *models.UserProfile, errVal error)
	InsertUser(userIn *pkg.UserIn) error
	DeleteUserByPubId(userPubId *uuid.UUID) error
	Migrate(dbName *string) (errVal error)
}
