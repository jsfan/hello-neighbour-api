package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/dal"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

type Store struct {
	DAL dal.AccessInterface
}

type DataInterface interface {
	GetDAL(ctx context.Context) (dalInstance dal.AccessInterface, commitFunc func() error, rollbackFunc func() error, errVal error)
	Migrate(dbName *string) (errVal error)
	GetUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error)
	RegisterUser(ctx context.Context, userIn *pkg.UserIn) (user *models.UserProfile, errVal error)
	DeleteUser(ctx context.Context, userPubId *uuid.UUID) error
	AddChurch(ctx context.Context, churchIn *pkg.ChurchIn) (church *models.ChurchProfile, errVal error)
	PromoteToLeader(ctx context.Context, userPubId *uuid.UUID, churchPubId *uuid.UUID) error
}
