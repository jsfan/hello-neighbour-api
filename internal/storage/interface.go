package storage

import (
	"context"
	"github.com/jsfan/hello-neighbour/internal/storage/dal"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

type Store struct {
	dal dal.AccessInterface
}

type DataInterface interface {
	GetDAL(ctx context.Context) (dalInstance dal.AccessInterface, commitFunc func() error, errVal error)
	Migrate(dbName *string) (errVal error)
	GetUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error)
	UserRegister(ctx context.Context, userIn *pkg.UserIn) (user *models.UserProfile, errVal error)
}
