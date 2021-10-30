package interfaces

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

// DbInteractionInterface provides an abstraction over both simple database connections and transactions
type DbInteractionInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// AccessInterface provides low level access to the database backend
type AccessInterface interface {
	Clone() AccessInterface

	BeginTransaction() error
	CancelTransaction() error
	CompleteTransaction() error

	DeleteUserByPubId(ctx context.Context, userPubId *uuid.UUID) error
	InsertChurch(ctx context.Context, churchIn *pkg.ChurchIn) (church *models.ChurchProfile, errVal error)
	InsertUser(ctx context.Context, userIn *pkg.UserIn) error
	MakeLeader(ctx context.Context, churchPubId *uuid.UUID, userPubId *uuid.UUID) error
	Migrate(dbName *string) (errVal error)
	SelectChurchByEmail(tx context.Context,email string) (church *models.ChurchProfile, errVal error)
	SelectUserByEmail(tx context.Context,email string) (user *models.UserProfile, errVal error)
	UpdateChurchActivationStatus(tx context.Context, churchPubId *uuid.UUID, isActive bool) error
}
