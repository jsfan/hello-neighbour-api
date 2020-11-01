package dal

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

type DAL struct {
	Db  *sql.DB
	ctx context.Context
	tx  *sql.Tx
}

type AccessInterface interface {
	SetupDal(ctx context.Context) (commit func() error, errVal error)
	SelectUserByEmail(email string) (user *models.UserProfile, errVal error)
	InsertUser(userIn *pkg.UserIn) error
	DeleteUserByPubId(userPubId *uuid.UUID) error
	Migrate(dbName *string) (errVal error)
	InsertChurch(churchIn *pkg.ChurchIn) error
	SelectChurchByEmail(email string) (church *models.ChurchProfile, errVal error)
}
